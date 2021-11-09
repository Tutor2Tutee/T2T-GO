package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email     string
	Nickname  string
	UserID    string
	ExpiresAt int64
	jwt.StandardClaims
}

func GenerateAccessToken(email string, nickname string, userID string) string {
	claims := &SignedDetails{
		Email:     email,
		Nickname:  nickname,
		UserID:    userID,
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return token
}

func GenerateRefreshToken(email string, nickname string, userID string) string {
	refreshClaims := &SignedDetails{
		Email:     email,
		Nickname:  nickname,
		UserID:    userID,
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(172)).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(172)).Unix(),
		},
	}

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return refreshToken
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

func VerifyUserAuthUsingJWT(c *gin.Context, userid, nickname, email string) bool {
	jwtUserEmail := c.Request.Header.Get("UserEmail")
	jwtUserID := c.Request.Header.Get("UserID")
	jwtUserNickname := c.Request.Header.Get("UserNickname")

	if jwtUserEmail != email || jwtUserID != userid || jwtUserNickname != nickname {
		return false
	}

	return true
}
