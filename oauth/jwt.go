package oauth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingSecret = []byte("my_secret")

// IssueToken creates JWT token
func IssueToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":         "bar",
		"issuingTime": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("my_secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeToken decodes a JWT token and returns its claims
func DecodeToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, err
}
