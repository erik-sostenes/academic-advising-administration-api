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
		Name                  string  `json:"name"`
		Email                 string  `json:"email"`
		IsActive              bool    `json:"is_active"`
		Cubicle               string  `json:"cubicle"`
		SubjectId             string  `subject_id`
		UniversityCourseId    string  `university_course_id`
		SubcoordinatorTuition string  `subcoordinator_tuition`
		CoordinatorTuition    string  `coordinator_tuition`
	}
	// TeacherCards collection of TeacherCard
	TeacherCards []TeacherCard

	// Login represents the access structure that a user must have
	Login struct {
		Tuition  string `json:"tuition"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Response map used for http response error
	Response map[string] interface{}
)


