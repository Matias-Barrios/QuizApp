package models

import "github.com/dgrijalva/jwt-go"

// User : Model struct for user data
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// Claim  :
type Claim struct {
	User User
	jwt.StandardClaims
}
