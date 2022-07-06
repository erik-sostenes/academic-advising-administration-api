package model

type(
	// Schedule represents a schedule
	Schedule struct {
		ScheduleId 	      string `json:"schedule_id"`
		ScheduleAt 	      string `json:"schedule_at"`
		FromDate 	        string `json:"from_date"`
		ToDate 		        string `json:"to_date"`
		StudentAccountant uint8  `json:"student_accountant"`
		TeacherTuition 	  string `json:"teacher_tuition"`
	}

	// TeacherSchedule represents a teacher schedule
	TeacherSchedule struct {
		Day            string   `json:"day"` 
		TeacherName    string   `json:"teacher_name"`
		SurnameTeacher string   `json:"surname"`
		Schedule       Schedule `json:"schedule"`
	}
	// TeacherSchedules collection of TeacherSchedule
	TeacherSchedules []TeacherSchedule 
	
	// TeacherCard represents a teacher card 	
	TeacherCard struct {
		Tuition               string  `json:"tuition"`
		Name                  string  `json:"name"`
		Surnames              string  `json:"surnames"`
		Email                 string  `json:"email"`
		IsActive              bool    `json:"is_active"`
		Cubicle               string  `json:"cubicle"`
		SubjectId             string  `json:"subject_id"`
		UniversityCourseId    string  `json:"university_course_id"`
		SubcoordinatorTuition string  `json:"subcoordinator_tuition"`
		CoordinatorTuition    string  `json:"coordinator_tuition"`
	}

	// TeacherCards collection of TeacherCard
	TeacherCards []TeacherCard

	// StudentRequest represents the requests of all the students who have asked him to take an advisory
	StudentRequest struct {
		Tuition           string `json:"tuition"`
		Name              string `json:"name"`
		Email             string `json:"email"`
		CubicleNumber     uint16 `json:"cubicle_number"`
		Subject           string `json:"subject"`
		BuildingNumber    string `json:"building_number"`
		AdvisoryId        string `json:"advisory_id"`
		TeacherScheduleId string `json:"teacher_schedule_id"`
	}
	// StudentRequests collection of StudentRequest
	StudentRequests []StudentRequest 
	
	StudentAcceptedRequest struct {
		Tuition          string `json:"tuition"`
		Name             string `json:"name"`
		Email            string `json:"email"`
		CubicleNumber    uint16 `json:"cubicle"`
		Subject          string `json:"subject"`
		UniversityCourse string `json:"university_course"`
	}
	//StudentAcceptedRequests
	StudentAcceptedRequests []StudentAcceptedRequest


	// Login represents the access structure that a user must have
	Login struct {
		Tuition  string `json:"tuition"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Authorization
	Authorization struct {
		Tuition  string `json:"tuition"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}
	// Response map used for http response error
	Response map[string] interface{}
)


