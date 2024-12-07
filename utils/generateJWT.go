package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// the jwt secret to generate token and verify it later
var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
