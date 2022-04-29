package model

// Schedule represents a schedule 
type Schedule struct {
	ScheduleId 	      string `json:"schedule_id"`
	ScheduleAt 	      string `json:"schedule_at"`
	FromDate 	        string `json:"from_date"`
	ToDate 		        string `json:"to_date"`
	StudentAccountant uint8  `json:"student_accountant"`
	TeacherTuition 	  string `json:"teacher_tuition"`
}

// TeacherSchedule represents a teacher schedule
type TeacherSchedule struct {
	Day            string   `json:"day"` 
	TeacherName    string   `json:"teacher_name"`
	SurnameTeacher string   `json:"surname"`
	Schedule       Schedule `json:"schedule"`
}

// TeacherSchedules 
type TeacherSchedules []TeacherSchedule 

// Response map used for http response in json format
type Response map[string] interface{}
