package database

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, email string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	insert, err := sqlConnection.Query(`INSERT INTO Users 
						 (name,password_encrypted,email, created_on,active) 
						 VALUES
						 (?,?,?,?,?)`, username, string(bytes), email, time.Now().UTC().Unix(), true)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil
}
