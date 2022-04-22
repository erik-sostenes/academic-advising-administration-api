package model

import (
	"strings"
	"time"
	"github.com/google/uuid"
)

// Declaring layout constant
const layout = "2006-Jan-02"

// MockSchedule represent a Schedule
type MockSchedule struct {
	scheduleId     string
	scheduleAt     time.Time
	fromDate       time.Time
	toDate 	       time.Time 
	teacherTuition string	
}

// NewMockSchedule returns an instance of MockSchedule if everything is correct
func NewMockSchedule(scheduleId string, scheduleAt, fromDate, toDate, teacherTuition string) (*MockSchedule, error) {
	uuid, err  := NewScheduleId(scheduleId)
	if err != nil {
		return &MockSchedule{}, err
	}

	scheduleAtV, err := NewTime(scheduleAt)
	if err != nil {
		return &MockSchedule{}, err
	}

	fromDateV, err := NewTime(fromDate)
	if err != nil {
		return &MockSchedule{}, err
	}

	toDateV, err := NewTime(toDate)
	if err != nil {
		return &MockSchedule{}, err
	}

	teacherTuitionV, err := NewTeacherTuition(teacherTuition)
	if err != nil {
		return &MockSchedule{}, err
	}
	return &MockSchedule{
		scheduleId: uuid.String(),
		scheduleAt: scheduleAtV,
		fromDate: fromDateV,
		toDate: toDateV,
		teacherTuition: teacherTuitionV,
	}, nil
}

var ErrInvalidScheduleId = StatusBadRequest("Invalid Schedule Id.")

// NewScheduleId returns a new value of uuid.UUID 
func NewScheduleId(scheduleId string) (value uuid.UUID, err error) {
	value, err = uuid.Parse(scheduleId)

	if err != nil {
		err = ErrInvalidScheduleId
		return 
	} 

	return 
}


var ErrInvalidTime = StatusBadRequest("Invalid Time.")

// NewTime returns a new value of time.Time
func NewTime(value string) (t time.Time, err error) {
	t, err = time.Parse(layout, value)

		if err != nil {
			err = ErrInvalidTime 
			return 
		}

	return 
}

var ErrInvalidTeacherTuition = StatusBadRequest("Invalid Teacher Tuition.");

// NewTeacherTuition returns a string if the teacherTuition is valid
func NewTeacherTuition(teacherTuition string) (value string, err error) {
	value = teacherTuition

	if strings.TrimSpace(value) == "" {
		err = ErrInvalidTeacherTuition
		return
	}

	return
}
