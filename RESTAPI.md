# API documentation

## APIs for user
1. `GET` : "/user/fetch" : Fetch user details 
2. `POST` : "/user/register" : User Registration

        `Content-Type :"application/json"
        body : 
            [
                {
                    "userid" : 1,
                    "username" : "XYZ",
                    "email" : "xyz@ymail.com",
                    "contactno" : "0123456789",
                    "password" : "abcd",
                    "userflag" : "T"   (set T or S as per trainer or student)
                }
            ]`
3. `GET` : "/user/login/{username}/{password}" : Login 

        `parameters: `
        `    - username = string(data type)
             - password = string`
4. `DELETE` : "/user/delete/{userid}" : User can delete their account

        `parameters: `  
        `   - userid = int(data type)`

## APIs for course
5. `POST` : "/course/new" : Create new course only by trainer

        `Content-Type :"application/json"`        
        `body :
            [
                {
                    "courseid" : 1,
                    "coursename" : "Go",
                    "description" : "Capture all concepts of golang",
                    "videourl" : "https://www.youtube.com/watch?v=Q0sKAMal4WQ",
                    "trainerid" : 1
                }
            ]`
            
6. `GET` : "/course/fetch"  : Fetch all courses only by trainer
7. `DELETE` : "/course/delete/{courseid}" : Delete particular course by particular trainer

        `parameters: ` 
         `   - courseid = int`
8. `PUT` : "/course/update/{courseid}" : Update course by particular trainer

        `Content-Type :"application/json"`
       ` parameters:  
            - courseid = int`
        `body :
            [
                {
                    "coursename" : "Go",
                    "description" : "Capture all concepts of golang",
                    "videourl" : "https://www.youtube.com/watch?v=Q0sKAMal4WQ"
                }  
            ]`  
9. `GET` : "/course/fetch/{courseid} : Fetch Particular course only by trainer

        `parameters:`  
         `       - courseid = int`

## APIs for student_course
10. `POST` : "/course/student/new" : Student can enrolled in particular course

        `Content-Type :"application/json"`
        `body :
            [
                {
                        "csid" : 1,
                        "courseid" : 1,
                        "studentid" : 1
                }
            ]`
11. `GET` : "/course/student/fetch/{studentname}" : Student fetch all course details if he/she is enrolled in that course otherwise fetch few details

        `parameters:  `
         `       - studentname = string`
12. `DELETE` : "/course/student/unenroll/{studentid}/{courseid}" : Student Unenroll from particular course

        `parameters:  `
        `        - studentid = int
                - courseid = int`
13. `PUT` : "/course/student/block/{studentid}/{courseid}" : Trainer can block student in particular course

            `parameters:  `
             `   - studentid = int
                - courseid = int`
 