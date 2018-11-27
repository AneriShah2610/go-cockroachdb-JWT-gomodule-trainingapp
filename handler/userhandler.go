package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
	"github.com/gorilla/mux"
)

var ctxt context.Context

func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crConn, err := driver.ConnectDb()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(r.Context(), "crConn", crConn)

		next.ServeHTTP(w, r.WithContext(ctxt))
	})
}

// RegisterUser create new user
func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	user := model.User{}
	_ = json.NewDecoder(request.Body).Decode(&user)
	crConn := ctxt.Value("crConn").(*driver.DB)
	if _, err := crConn.DatabaseConn.Exec("INSERT INTO training.user_mst VALUES ($1,$2,$3,$4,$5,$6)", user.UserID, user.UserName, user.Email, user.Contactno, user.Password, user.UserFlag); err != nil {
		log.Fatal("Error while inserting data :", err)
	}
	fmt.Fprintf(writer, `Registration succssfully`)
}

/*func LogIn(writer http.ResponseWriter, request *http.Request) {
	requestUser := new(model.User)
	json.NewDecoder(request.Body).Decode(&requestUser)
	responseStatus, token := services.LogIn(requestUser)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(responseStatus)
	fmt.Fprintf(writer, token)
}*/

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
func BlockUser(writer http.ResponseWriter, request *http.Request) {
	user := model.User{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	if _, err := crConn.DatabaseConn.Exec("UPDATE training.course_mst SET (csflag=$1) WHERE courseid=$2", user.UserFlag, params["id"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `user updated successfully`)
	}
}
