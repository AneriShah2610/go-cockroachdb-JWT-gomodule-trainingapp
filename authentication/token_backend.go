package authentication

import (
	"log"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var secretkey = []byte("secret_key")
var user model.User

// GenerateToekn generate token
func GenerateToekn() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"userid":    user.UserID,
		"username":  user.UserName,
		"email":     user.Email,
		"contactno": user.Contactno,
		"password":  user.Password,
	}
	tokenstring, err := token.SignedString(secretkey)
	if err != nil {
		log.Fatal("Error while generating token ", err)
	}
	return tokenstring, err
}