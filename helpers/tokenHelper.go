package helpers

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email    string
	Nickname string
	UserID   string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, nickname string, userID string) string {
	claims := &SignedDetails{
		Email:    email,
		Nickname: nickname,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return token
}
