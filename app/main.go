package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Connect to database

	// Router Init
	router := mux.NewRouter()

	router.Use(handler.MiddleWareHandler)
	// Functions for user
	router.HandleFunc("/user/fetch", handler.FetchUser).Methods("GET")
	router.HandleFunc("/user/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/user/login/{username}/{password}", handler.ReadData(handler.LogIn)).Methods("GET")
	router.HandleFunc("/user/delete/{userid}", handler.DeleteUser).Methods("DELETE")

	// Functions on course by trainer
	router.HandleFunc("/course/new", handler.CreateCourse).Methods("POST")
	router.HandleFunc("/course/fetch", handler.FetchCourse).Methods("GET")
	router.HandleFunc("/course/delete/{courseid}", handler.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/course/update/{courseid}", handler.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/fetch/{courseid}", handler.FetchParticularCourse).Methods("GET")

	// Functions for studentcourse
	router.HandleFunc("/course/student/new", handler.EnrollInCourse).Methods("POST")
	router.HandleFunc("/course/student/fetch/{studentname}", handler.FetchEnrolledCourse).Methods("GET")
	router.HandleFunc("/course/student/unenroll/{studentid}/{courseid}", handler.UnEnroll).Methods("DELETE")
	router.HandleFunc("/course/student/block/{studentid}/{courseid}", handler.BlockUser).Methods("PUT")

	log.Fatal("Error while routing handler", http.ListenAndServe(":8000", router))
}
