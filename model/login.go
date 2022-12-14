package model

import "github.com/dgrijalva/jwt-go"

// Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim contendrĂ¡ el cuerpo que tendrĂ¡ el token
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
