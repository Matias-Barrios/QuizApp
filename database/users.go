package database

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Matias-Barrios/QuizApp/models"
	"golang.org/x/crypto/bcrypt"
)

// GetUser :
func GetUser(password, email string) (models.User, error) {
	email = strings.ToLower(email)
	var user models.User
	err := sqlConnection.QueryRow(`
		SELECT id,name, email, password_encrypted
		FROM Users U
		WHERE active = true
		AND email = ?
		AND id NOT IN (
			SELECT DISTINCT user_id FROM Unsuccessful_login_attempts
					WHERE Unsuccessful_login_attempts.user_id = U.id 
					AND 
					(
						SELECT COUNT(*) FROM Unsuccessful_login_attempts ula
						WHERE ula.attempted_on > ?
					) >= 3
				)
	`, email, time.Now().Add(-10*time.Minute).UTC().Unix()).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err.Error())
		return models.User{}, err
	}
	if CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	_, err = sqlConnection.Exec(`
		INSERT INTO Unsuccessful_login_attempts (user_id, attempted_on)
		VALUES((SELECT id FROM Users WHERE email = ?),?)
	`, email, time.Now().UTC().Unix())

	if err != nil {
		log.Println(err.Error())
	}
	return models.User{}, fmt.Errorf("Unauthorized\n")
}

// CreateUser :
func CreateUser(username, password, email string) error {
	email = strings.ToLower(email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = sqlConnection.Exec(`INSERT INTO Users 
						 (name,password_encrypted,email, created_on,active) 
						 VALUES
						 (?,?,?,?,?)`, username, string(bytes), email, time.Now().UTC().Unix(), true)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// UpdateUserPassword :
func UpdateUserPassword(id int, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = sqlConnection.Exec(`
								UPDATE Users
								SET password_encrypted = ?
								WHERE id = ? 
								`, string(bytes), id)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func SetNewPassword(email, password string) error {
	email = strings.ToLower(email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = sqlConnection.Exec(`
	DELIMITER $
    BEGIN NOT ATOMIC
		IF (SELECT count(*) FROM Users WHERE email = ? ) != 1 
		THEN 
		    SIGNAL SQLSTATE '12345' SET MESSAGE_TEXT = 'Unauthorized';
		ELSE
			UPDATE Users 
			SET password_encrypted = ? 
			WHERE email = ?;
		END IF
	END $
	DELIMITER ;`, email, string(bytes), email)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
