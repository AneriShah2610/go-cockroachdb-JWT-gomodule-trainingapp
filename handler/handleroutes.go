package handler

import (
	"net/http"

	repository "github.com/s1s1ty/go-mysql-crud/repository"
)

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

type Post struct {
	repo repository.PostRepo
}

// FetchUserDetails fetch all the details about user
func FetchUserDetails(w http.ResponseWriter, r *http.Request) {

}

// UserGetByID to get particular user detail by userid
func UserGetByID(w http.ResponseWriter, r *http.Request) {

}

// CreateUser to create/insert new data of user
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser to delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser to update user detail by userid
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// FetchCourseDetails to fetch all course details
func FetchCourseDetails(w http.ResponseWriter, r *http.Request) {

}

// CourseDetailByID to get particular course detail by courseid
func CourseDetailByID(w http.ResponseWriter, r *http.Request) {

}

// CreateCourse to insert new course details only by trainer
func CreateCourse(w http.ResponseWriter, r *http.Request) {

}

// DeleteCourse to delete course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {

}

// UpdateCourse to update course details
func UpdateCourse(w http.ResponseWriter, r *http.Request) {

}

// GetTrainerBycourse to get coursedetail by trainer id
func GetTrainerBycourse(w http.ResponseWriter, r *http.Request) {

}

// GetStudentByCourse to get courses as per student id
func GetStudentByCourse(w http.ResponseWriter, r *http.Request) {

}
