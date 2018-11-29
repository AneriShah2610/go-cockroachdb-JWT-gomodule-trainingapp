package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/authentication"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
	"github.com/gorilla/mux"
)

// SessionTokenResponse token
type SessionTokenResponse struct {
	Token string `json:"sessionToken"`
}

var ctxt context.Context

// MiddleWareHandler  middleware
func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn, err := driver.ConnectDb()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(ctxt))
	})
}

// FetchUser to retriev all data from user table
func FetchUser(writer http.ResponseWriter, request *http.Request) {
	user := model.User{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	rows, err := crConn.DatabaseConn.Query("SELECT * FROM training.user_mst")
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	fmt.Fprintf(writer, `All data about users`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.Contactno, &user.Password, &user.UserFlag); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		json.NewEncoder(writer).Encode(user)
	}
}

/* ----- registration ----- */

// RegisterUser for user registration
func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	/* logic for check existing user data
	user, err := readUser() // Call readUser to get user email from user table to
	if err != nil {
		log.Fatal("Error while calling readUser function", err)
	}
	fmt.Println(user)*/
	tokenString, err := authentication.GenerateToken() // Generate token
	if err != nil {
		log.Fatal("Error at generating token : ", err)
	}
	InsertData(writer, request, SessionTokenResponse{Token: tokenString})
}

// InsertData function to insert data into user table
func InsertData(writer http.ResponseWriter, request *http.Request, response interface{}) {
	users := model.User{}
	_ = json.NewDecoder(request.Body).Decode(&users)
	crConn := ctxt.Value("crConn").(*driver.DB)
	if _, err := crConn.DatabaseConn.Exec("INSERT INTO training.user_mst VALUES ($1,$2,$3,$4,$5,$6)", users.UserID, users.UserName, users.Email, users.Contactno, users.Password, users.UserFlag); err != nil {
		log.Fatal("Error while inserting data :", err)
	}
	fmt.Fprintf(writer, `user created successfully`)
}

/* ----- Complete registration -----*/

/* ----- Login ----- */

// ReadData middleware
func ReadData(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users := model.User{}
		_ = json.NewDecoder(request.Body).Decode(&users)
		params := mux.Vars(request)
		crConn := ctxt.Value("crConn").(*driver.DB)
		result := crConn.DatabaseConn.QueryRow("SELECT username,user_password,email FROM training.user_mst WHERE username=$1", params["username"])
		err := result.Scan(&users.UserName, &users.Password, &user.Email)
		if err != nil {
			log.Fatal("Error while scanning password", err)
			return
		}
		if users.UserName != params["username"] || users.Password != params["password"] {
			fmt.Fprintf(writer, `Invalid user or pasword`)
			return
		}
		next(writer, request)
	}
}

// LogIn function
func LogIn(writer http.ResponseWriter, request *http.Request) {
	tokenString, err := authentication.GenerateToken() // Generate token
	if err != nil {
		log.Fatal("Error at generating token : ", err)
	}
	writeLogIn(writer, request, SessionTokenResponse{Token: tokenString})
}
func writeLogIn(writer http.ResponseWriter, request *http.Request, response interface{}) {
	fmt.Fprintf(writer, `Login successfull`)
}

/* ----- Complete Login ----- */

/* ----- Delete user ----- */

// DeleteUser to delete particular student
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	if _, err := crConn.DatabaseConn.Exec("DELETE FROM training.user_mst WHERE userid=$1", params["userid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `user deleted successfully`)
	}
}

/* ----- Complete Delete user ----- */
