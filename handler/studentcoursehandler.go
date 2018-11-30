package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/driver"
	"github.com/AneriShah2610/go-cockroachdb-JWT-gomodule-trainingapp/model"
)

var (
	course        = model.Course{}
	studentcourse = model.CourseStudent{}
	user          = model.User{}
)

// EnrollInCourse student enroll in particular course
func EnrollInCourse(writer http.ResponseWriter, request *http.Request) {
	_ = json.NewDecoder(request.Body).Decode(&studentcourse)
	crConn := ctxt.Value("crConn").(*driver.DB)
	if _, err := crConn.DatabaseConn.Exec("INSERT INTO training.studentcourse VALUES ($1,$2,$3)", studentcourse.CourseStudentID, studentcourse.CourseID, studentcourse.StudentID); err != nil {
		log.Fatal("Error while inserting data :", err)
	}
	fmt.Fprintf(writer, `course enollment successfully`)
}

// FetchEnrolledCourse for student can fetch own enrolled student
func FetchEnrolledCourse(writer http.ResponseWriter, request *http.Request) {
	// To do: check Loeedin or not
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	result := crConn.DatabaseConn.QueryRow("SELECT studentcourse.studentid,studentcourse.csflag FROM training.studentcourse WHERE studentid=$1", params["userid"])
	err := result.Scan(&studentcourse.StudentID, &studentcourse.CsFlag)
	if err != nil {
		log.Fatal("Error while fetching student id and csflag", err)
	}
	if studentcourse.StudentID == 0 {
		fmt.Fprintf(writer, `You have no enrolled courses`)
	} else {
		if studentcourse.CsFlag == "0" {
			row, err := crConn.DatabaseConn.Query("SELECT course_mst.coursename,course_mst.description,course_mst.videourl,course_mst.trainerid FROM training.course_mst,training.studentcourse,training.user_mst WHERE (studentcourse.courseid=course_mst.courseid) and (studentcourse.studentid=user_mst.userid) and studentid=$1", params["userid"])
			if err != nil {
				log.Fatal("Error while fetching course details", err)
			}
			for row.Next() {
				if err := row.Scan(&course.CourseName, &course.Description, &course.VideoURL, &course.TrainerID); err != nil {
					log.Fatal("Scannin data", err)
				}
				fmt.Fprintf(writer, course.CourseName, course.Description, course.VideoURL, course.TrainerID)
			}
		} else {
			row, err := crConn.DatabaseConn.Query("SELECT course_mst.coursename,course_mst.description FROM training.course_mst,training.user_mst,training.studentcourse WHERE (studentcourse.courseid=course_mst.courseid) and (studentcourse.studentid=user_mst.userid) and studentid=$1", params["userid"])
			if err != nil {
				log.Fatal("Error while fetching course details", err)
			}
			for row.Next() {
				if err := row.Scan(&course.CourseName, &course.Description); err != nil {
					log.Fatal("Scannin data", err)
				}
				fmt.Fprintf(writer, course.CourseName, course.Description)
			}
		}
	}
}

// UnEnroll to unenroll from course
func UnEnroll(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	if _, err := crConn.DatabaseConn.Exec("DELETE FROM training.studentcourse WHERE studentid=$1 and courseid=$2", params["studentid"], params["courseid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `course unenroll successfully`)
	}
}

// BlockUser to block particular student by trainer
func BlockUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	_ = json.NewDecoder(request.Body).Decode(&studentcourse)
	if _, err := crConn.DatabaseConn.Exec("UPDATE training.studentcourse SET csflag = '1' WHERE studentid = $1 and courseid = $2", params["studentid"], params["courseid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `block student successfully`)
	}
}
