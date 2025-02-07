package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, name string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"name":    name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateToken(tokenString string, jwtSecret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if len(jwtSecret) == 0 {
			return nil, errors.New("JWT secret key is missing")
		}
		return []byte(jwtSecret), nil
	})
}
