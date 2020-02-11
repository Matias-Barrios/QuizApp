package models

import "github.com/dgrijalva/jwt-go"

// User : Model struct for user data
type User struct {
	Name string
}

// Claim  :
type Claim struct {
	Username string
	jwt.StandardClaims
}
