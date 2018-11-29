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
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	result := crConn.DatabaseConn.QueryRow("select user_mst.username,studentcourse.studentid from training.user_mst,training.course_mst,training.studentcourse where user_mst.username=$1 and studentcourse.courseid=course_mst.courseid and studentcourse.studentid=user_mst.userid", params["studentname"])
	err := result.Scan(&user.UserName, &studentcourse.StudentID)
	if err != nil {
		log.Fatal("Error while scannig that student enrolled or not", err)
	}
	studentid := studentcourse.StudentID
	// if username is exist in course enrollment than they can get all the course details
	if user.UserName != " " {
		rows, err := crConn.DatabaseConn.Query("select course_mst.coursename,course_mst.description,course_mst.videourl from training.course_mst,training.studentcourse where (course_mst.courseid = studentcourse.courseid) and (studentcourse.studentid=$1) and studentcourse.csflag='0' ", studentid)
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
	} else {
		rows, err := crConn.DatabaseConn.Query("select course_mst.coursename,course_mst.description from training.course_mst,training.studentcourse where (course_mst.courseid = studentcourse.courseid) and (studentcourse.studentid=$1) and studentcourse.csflag='0' ", studentid)
		if err != nil {
			log.Fatal("Error while retrieving data", err)
		}
		fmt.Fprintf(writer, `Course details`)
		fmt.Fprintf(writer, "\n")
		for rows.Next() {
			if err := rows.Scan(&course.CourseName, &course.Description); err != nil {
				log.Fatal("Error while scanning data", err)
			}
			fmt.Fprintf(writer, "\n", course.CourseName, course.Description)
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
	checkuser, err := checkUser(request)
	if err != nil {
		log.Fatal("cana not find u as a trainer", err)
	}
	if checkuser == "T" {
		if _, err := crConn.DatabaseConn.Exec("UPDATE training.studentcourse SET csflag = '1' WHERE studentid = $1 and courseid = $2", params["studentid"], params["courseid"]); err != nil {
			log.Fatal("Error while deleting course", err)
		} else {
			fmt.Fprintf(writer, `block student successfully`)
		}
	} else {
		fmt.Fprintf(writer, `You are not authorized person`)
	}

}
func checkUser(request *http.Request) (string, error) {
	params := mux.Vars(request)
	crConn := ctxt.Value("crConn").(*driver.DB)
	row := crConn.DatabaseConn.QueryRow("select user_mst.user_flag from training.user_mst,training.course_mst where (user_mst.userid=course_mst.trainerid) and courseid=$1", params["courseid"])
	err := row.Scan(&user.UserFlag)
	if err != nil {
		log.Fatal("Error", err)
	}
	return user.UserFlag, err
}
