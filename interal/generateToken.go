package interal

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

//todo:add .evn file

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
	//fmt.Println(tokenString)
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, string, string) {
	fmt.Println(tokenString)
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
	fmt.Println(login, role)
	return true, login, role
	//return true, "login, role", ""
}
