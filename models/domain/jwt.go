package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
