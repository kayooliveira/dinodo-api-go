package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	jwt.RegisteredClaims
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

var tokenSecret = os.Getenv("JWT_SECRET")

func GenerateUserToken(id string, name string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenPayload{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		ID:    id,
		Name:  name,
		Email: email,
	})

	signedString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ParseUserToken(token string) (*TokenPayload, error) {

	var tokenPayload TokenPayload

	parsed, err := jwt.ParseWithClaims(token, &tokenPayload, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return &tokenPayload, nil
}
