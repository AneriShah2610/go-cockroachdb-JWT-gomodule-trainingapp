package authentication

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

func verifyToken(tokenString string) (jwt.Claims, error) {
	var signedKey = []byte("secretSignedKey")
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		log.Fatal("Error while verify token : ", err)
	}
	return token.Claims, err
}
