package main

import (
	"fmt"
	"log"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/authentication"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Connect to database
	connection, err := driver.ConnectDb()
	if err != nil {
		log.Fatal("Error at db connection ", err)
	}
	defer connection.DatabaseConn.Close()
	fmt.Println(connection, err)

	// Router Init
	router := mux.NewRouter()
	router.Use(authentication.authenticateJWT)
	handleFunc := handler.NewPostHandler(connection)
	router.HandleFunc("/user",handler.FetchUserDetails).Methods("GET")
	router.HandleFunc("/user/{id}",handler.UserGetByID).Methods("GET")
	router.HandleFunc("/user",handler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}",handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}",handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/course",handler.FetchCourseDetails).Methods("GET")
	router.HandleFunc("/course/{id}",handler.CourseDetailByID).Methods("GET")
	router.HandleFunc("/course",handler.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}",handler.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/course/{id}",handler.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course_trainer/{id}",handler.GetTrainerBycourse).Methods("GET")
	router.HandleFunc("/course_student/{id}",handler.GetStudentByCourse).Methods("GET")
	log.Fatal("Error while routing :",http.ListenAndServe(":8000",router))
}
