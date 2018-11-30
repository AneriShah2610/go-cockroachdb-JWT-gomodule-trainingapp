# Tables

    ### user_mst
        -
        userid  int PK 
        username string 
        email string 
        contactno string 
        user_password string
        user_flag string

    ### course_mst
        -
        courseid PK int 
        coursename string
        description string
        videourl string
        trainerid int FK >- user_mst.userid
    
    ### studentcourse
        - 
        csid int PK
        courseid int FK >- course_mst.courseid
        studentid int FK >- user_mst.userid
        csflag string (By default set 0 i.e unblock)