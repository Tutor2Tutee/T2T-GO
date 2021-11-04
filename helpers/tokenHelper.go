package helpers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email     string
	Nickname  string
	UserID    string
	ExpiresIn int64
	jwt.StandardClaims
}

func GenerateAllTokens(email string, nickname string, userID string) (string, string) {
	claims := &SignedDetails{
		Email:     email,
		Nickname:  nickname,
		UserID:    userID,
		ExpiresIn: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		Email:     email,
		Nickname:  nickname,
		UserID:    userID,
		ExpiresIn: time.Now().Local().Add(time.Hour * time.Duration(172)).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(172)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return token, refreshToken
}

func VerifyToken(token string) (*jwt.Token, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	return t, err
}
