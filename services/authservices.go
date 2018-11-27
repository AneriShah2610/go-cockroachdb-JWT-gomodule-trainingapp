package services

import (
	"fmt"
	"net/http"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/authentication"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
)

func LogIn(requestUser *model.User) (int, string) {
	token, err := authentication.GenerateToekn()
	if err != nil {
		return http.StatusInternalServerError, ""
	} else {
		return http.StatusOK, ""
	}
	fmt.Println(token)
	return http.StatusUnauthorized, token
}
