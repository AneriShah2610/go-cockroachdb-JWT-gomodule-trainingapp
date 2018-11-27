package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
	"github.com/gorilla/mux"
)

func CreateCourse(writer http.ResponseWriter, request *http.Request) {
	course := model.Course{}
	_ = json.NewDecoder(request.Body).Decode(&course)
	crConn := ctxt.Value("crConn").(*driver.DB)
	if _, err := crConn.DatabaseConn.Exec("INSERT INTO training.course_mst VALUES ($1,$2,$3,$4,$5)", course.CourseID, course.CourseName, course.Description, course.VideoURL, course.TrainerID); err != nil {
		log.Fatal("Error while inserting data :", err)
	}
	fmt.Fprintf(writer, `course created successfully`)
}
func FetchCourse(writer http.ResponseWriter, request *http.Request) {
	course := model.Course{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	rows, err := crConn.DatabaseConn.Query("SELECT * FROM training.course_mst")
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	fmt.Fprintf(writer, `All data about course`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&course.CourseID, &course.CourseName, &course.Description, &course.VideoURL, &course.TrainerID); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		json.NewEncoder(writer).Encode(course)
	}
}
func DeleteCourse(writer http.ResponseWriter, request *http.Request) {

	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	if _, err := crConn.DatabaseConn.Exec("DELETE FROM training.course_mst WHERE courseid=$1", params["courseid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `course deleted successfully`)
	}
}
func UpdateCourse(writer http.ResponseWriter, request *http.Request) {
	course := model.Course{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	if _, err := crConn.DatabaseConn.Exec("UPDATE training.course_mst SET (courseid=$1,description=$2,videourl=$3) WHERE courseid = $4", course.CourseName, course.Description, course.VideoURL, params["courseid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `course updated  successfully`)
	}
}
func FetchParticularCourse(writer http.ResponseWriter, request *http.Request) {
	course := model.Course{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	row, err := crConn.DatabaseConn.Query("SELECT * FROM training.course_mst WHERE courseid=$1", params["courseid"])
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	for row.Next() {
		if err := row.Scan(&course.CourseID, &course.CourseName, &course.Description, &course.VideoURL, &course.TrainerID); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		json.NewEncoder(writer).Encode(course)
	}
}
