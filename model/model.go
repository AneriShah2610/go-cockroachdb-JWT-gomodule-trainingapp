package model

// User model
type User struct {
	UserID    int    `json:"userid"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Contactno string `json:"contactno"`
	Password  string `json:"password"`
	UserFlag  string `json:"userflag"`
}

// Course model
type Course struct {
	CourseID    int    `json:"courseid"`
	CourseName  string `json:"coursename"`
	Description string `json:"description"`
	VideoURL    string `json:"videourl"`
	TrainerID   int    `json:"trainerid"`
}

// CourseStudent model
type CourseStudent struct {
	CourseStudentID int `json:"csid"`
	StudentID       int `json:"studentid"`
	CourseDetail    *Course
}
