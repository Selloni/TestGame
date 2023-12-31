package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const swToken = "grandpat"

func GenerateToken(login, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	secret := []byte(swToken)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("Token error - %v", err)
		return "", err
	}
	log.Println("create token with role", role)
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, string, string) {
	secret := []byte(swToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return false, "", ""
	}

	claims := token.Claims.(jwt.MapClaims)
	login := claims["login"].(string)
	role := claims["role"].(string)
	return true, login, role
}
