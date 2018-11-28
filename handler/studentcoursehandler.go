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

// EnrollInCourse student enroll in particular course
func EnrollInCourse(writer http.ResponseWriter, request *http.Request) {
	studentcourse := model.CourseStudent{}
	_ = json.NewDecoder(request.Body).Decode(&studentcourse)
	crConn := ctxt.Value("crConn").(*driver.DB)
	var csflag = "0"
	if _, err := crConn.DatabaseConn.Exec("INSERT INTO training.studentcourse VALUES ($1,$2,$3,$4)", studentcourse.CourseStudentID, studentcourse.CourseID, studentcourse.StudentID, csflag); err != nil {
		log.Fatal("Error while inserting data :", err)
	}
	fmt.Fprintf(writer, `course enollment successfully`)
}

// FetchCourseByStudent for fetch course by student
func FetchCourseByStudent(writer http.ResponseWriter, request *http.Request) {
	course := model.Course{}
	user := model.User{}
	crConn := ctxt.Value("crConn").(*driver.DB)
	rows, err := crConn.DatabaseConn.Query("select course_mst.coursename,course_mst.description,course_mst.videourl,user_mst.username from training.course_mst,training.user_mst where user_mst.user_flag='T' and user_mst.userid=course_mst.trainerid;")
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	fmt.Fprintf(writer, `All data about course`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&course.CourseName, &course.Description, &course.VideoURL, &user.UserName); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		fmt.Fprintf(writer, "\n", course.CourseName, course.Description, course.VideoURL, user.UserName)
	}
}

// FetchEnrolledCourse for student can fetch own enrolled student
func FetchEnrolledCourse(writer http.ResponseWriter, request *http.Request) {
	//studentcourse := model.CourseStudent{}
	course := model.Course{}
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	rows, err := crConn.DatabaseConn.Query("select course_mst.coursename,course_mst.description,course_mst.videourl from training.course_mst,training.studentcourse where (course_mst.courseid = studentcourse.courseid) and (studentcourse.studentid=$1)", params["studentid"])
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	fmt.Fprintf(writer, `My Enrolled courses`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&course.CourseName, &course.Description, &course.VideoURL); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		fmt.Fprintf(writer, "\n", course.CourseName, course.Description, course.VideoURL)
	}
}

// UnEnroll to unenroll from course
func UnEnroll(writer http.ResponseWriter, request *http.Request) {
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	if _, err := crConn.DatabaseConn.Exec("DELETE FROM training.studentcourse WHERE studentid=$1 and courseid=$2", params["studentid"], params["courseid"]); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `course unenroll successfully`)
	}
}

// BlockUser to block particular student by trainer
func BlockUser(writer http.ResponseWriter, request *http.Request) {
	studentcourse := model.CourseStudent{}
	_ = json.NewDecoder(request.Body).Decode(&studentcourse)
	crConn := ctxt.Value("crConn").(*driver.DB)
	params := mux.Vars(request)
	var csflag = "1"
	if _, err := crConn.DatabaseConn.Exec("UPDATE training.studentcourse SET csflag = $3 WHERE studentid = $1 and courseid = $2", params["studentid"], params["courseid"], csflag); err != nil {
		log.Fatal("Error while deleting course", err)
	} else {
		fmt.Fprintf(writer, `course updated  successfully`)
	}
}
