package authentication

import (
	"fmt"
	"log"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var user model.User
var signedKey = []byte("secretSignedKey")

func authenticateJWT() {
	tokenString := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userid":    user.UserName,
		"username":  user.Password,
		"email":     user.Email,
		"contactno": user.Contactno,
		"password":  user.Password,
		"userflag":  user.UserFlag,
	})
	tokenVerify, err := tokenString.SignedString(signedKey)
	if err != nil {
		log.Fatal("Error while verify the token", err)
	}
	claims, err := verifyToken(tokenVerify)
	fmt.Println(claims)
}
