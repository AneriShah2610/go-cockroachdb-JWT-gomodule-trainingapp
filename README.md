# go-cockroachdb-JWT-gomodule-trainingapp
 A simple training app using golang with help of JWT , cockroach db and go module

## Prerequisite
1. [Install golang](https://golang.org/dl/)
2. Setup GOPATH [Link1](https://golang.org/doc/code.html#GOPATH) and [Link2](https://github.com/golang/go/wiki/GOPATH)
3. [Install cockroachdb](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-windows.html)
4. Check Binary file of cockroachdb : `.\cockroach.exe` run in cmd
5. Start node in cockroachdb : `cockroach start --insecure` run in cmd where cockroachdb is intalled
6. Test cluster : `cockroach sql --insecure` run in cmd where cockroachdb is intalled

### Package used
1. github.com/dgrijalva/jwt-go = For token authentication
2. github.com/lib/pq = For cokcroachdb 
3. github.com/gorilla/mux = As a router to carry handler requests

#### APIs (Operations for APIs)

- `GET` : [/user/fetch] - Retrieve all users data
- `POST` : [/user/register] - Registeration (Two types of users : Student & Trainer)
- `GET` : [/user/login/{username}/{password}] - Login
- `DELETE` : [/user/delete/{userid}] - Delete user
- `POST` : [/course/new] - Create new course only by trainer
- `GET` : [/course/fetch] - Fetch all courses only by trainer
- `DELETE` : [/course/delete/{courseid}] - Delete course by trainer
- `PUT` : [/course/update/{courseid}] - Update course by trainer
- `GET` : [/course/fetch/{courseid}] - Fetch particular course by courseid
- `POST` : [/course/student/new] - Student enroll in course
- `GET` : [/course/student/fetch/{studentname}] - Student fetch all course details if he/she is enrolled in that course otherwise fetch few details
- `DELETE` : [/course/student/unenroll/{studentid}/{courseid}] - Student Unenroll from particular course
- `PUT` : [/course/student/block/{studentid}/{courseid}] - Trainer can block particular student for particular course 

##### Getting started

1. Clone repo
2. Run `go run main.go`
3. Run 'http://localhost:8000/' on postman or ant other client tool