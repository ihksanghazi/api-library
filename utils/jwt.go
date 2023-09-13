package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihksanghazi/api-library/models/domain"
)

func GenerateToken(req *domain.User, signed string, time time.Time) (string, error) {
	claims := domain.JWTClaims{
		Username: req.Username,
		Email:    req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time),
			ID:        req.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(signed))
	return result, err
}

func ParsingToken(yourToken string, signed string) (claims *domain.JWTClaims, err error) {

	token, errParsing := jwt.ParseWithClaims(yourToken, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signed), nil
	})

	if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errParsing
	}
}
