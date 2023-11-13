package interal

import "github.com/dgrijalva/jwt-go"

//todo:add .evn file

func GenerateToken(login, role string) string {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["role"] = role

	secret := []byte("nowy!sekrret-1213@KEYlis")
	tokenString, _ := token.SignedString(secret)

	return tokenString
}

func ValidateToken(tokenString string) (bool, string, string) {
	secret := []byte("nowy!sekrret-1213@KEYlis") // Замените на свой секретный ключ
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
