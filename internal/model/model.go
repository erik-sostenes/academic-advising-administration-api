package model

import (
	"time"
)

// Schedule represents a schedule 
type Schedule struct {
	ScheduleId 	  string    `json:"schedule_id"`
	ScheduleAt 	  time.Time `json:"schedule_at"`
	FromDate 	  time.Time `json:"from_date"`
	ToDate 		  time.Time `json:"to_date"`
	StudentAccountant uint8     `json:"student_accountant"`
	TeacherTuition 	  string    `json:"teacher_tuition"`
}

// TeacherSchedule represents a teacher schedule
type TeacherSchedule struct {
	Day            string   `json:"day"` 
	TeacherName    string   `json:"teacher_name"`
	SurnameTeacher string   `json:"surname"`
	Schedule       Schedule `json:"schedule"`
}
