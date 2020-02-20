package database

import (
	"fmt"
	"time"

	"github.com/Matias-Barrios/QuizApp/models"
	"golang.org/x/crypto/bcrypt"
)

// GetUser :
func GetUser(password, email string) (models.User, error) {
	var user models.User
	err := sqlConnection.QueryRow(`
		SELECT id,name, email, password_encrypted
		FROM Users
		WHERE active = true; 
	`).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	if CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return models.User{}, fmt.Errorf("Unauthorized\n")
}

// CreateUser :
func CreateUser(username, password, email string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	_, err = sqlConnection.Exec(`INSERT INTO Users 
						 (name,password_encrypted,email, created_on,active) 
						 VALUES
						 (?,?,?,?,?)`, username, string(bytes), email, time.Now().UTC().Unix(), true)

	if err != nil {
		return err
	}
	return nil
}
