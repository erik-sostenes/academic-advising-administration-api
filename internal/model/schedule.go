package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// ErrInvalidScheduleId error schedule id is invalid
var ErrInvalidScheduleId = StatusBadRequest("Invalid Schedule Id.")
// ErrInvalidTime error date parse occurs
var ErrInvalidTime = StatusBadRequest("Invalid Time.")
// ErrInvalidTeacherTuition error Teacher Tuition 
var ErrInvalidTeacherTuition = StatusBadRequest("Invalid Teacher Tuition.");


// Declaring layout constant
const layout = "2006-01-02 15:04:05"


// MockSchedule represents the mock of a schedule
type MockSchedule struct {
	scheduleId        string
	scheduleAt        string
	fromDate          string
	toDate 	          string 
	studentAccountant uint8
	teacherTuition    string	
}

// NewMockSchedule returns an instance of MockSchedule if everything is correct
func NewMockSchedule(scheduleId, scheduleAt, fromDate, toDate, teacherTuition string, studentAccountant uint8) (MockSchedule, error) {
	uuid, err  := newScheduleId(scheduleId)
	if err != nil {
		return MockSchedule{}, err
	}

	scheduleAtV, err := newTime(scheduleAt)
	if err != nil {
		return MockSchedule{}, err
	}

	fromDateV, err := newTime(fromDate)
	if err != nil {
		return MockSchedule{}, err
	}

	toDateV, err := newTime(toDate)
	if err != nil {
		return MockSchedule{}, err
	}

	teacherTuitionV, err := newTeacherTuition(teacherTuition)
	if err != nil {
		return MockSchedule{}, err
	}
	return MockSchedule{
		scheduleId: uuid,
		scheduleAt: scheduleAtV,
		fromDate: fromDateV,
		toDate: toDateV,
		studentAccountant: studentAccountant,
		teacherTuition: teacherTuitionV,
	}, nil
}
// ScheduleId represents the teacher schedule unique identifier
func (m *MockSchedule) ScheduleId() string {
	return m.scheduleId
}
// ScheduleAt represents the date the teacher's schedule was created
func (m *MockSchedule) ScheduleAt() string {
	return m.scheduleAt
}
// FromDate represents the start date of the academic advisory
func (m *MockSchedule) FromDate() string {
	return m.fromDate
}
// ToDate represents the end date of the academic advisory
func (m *MockSchedule) ToDate() string {
	return m.toDate 
}
// TeacherTuition represents the teacher's unique tuition
func (m *MockSchedule) TeacherTuition() string {
	return m.teacherTuition
}
// StudentAccountant represents the total number of students who are taking academic advice on that date 
func (m *MockSchedule) StudentAccountant() uint8 {
	return m.studentAccountant
}


// newScheduleId returns a new value of uuid.UUID 
func newScheduleId(scheduleId string) (string, error) {
	v, err := uuid.Parse(scheduleId)

	if err != nil {
		return "", ErrInvalidScheduleId
	} 

	return v.String(), nil
}


// newTime validates that a string has the correct format of a date 
func newTime(value string) (string, error) {
	_, err := time.Parse(layout, value)

		if err != nil {
			return "", ErrInvalidTime 
		}

		return  value, nil 
}

// newTeacherTuition returns a string if the teacherTuition is valid
func newTeacherTuition(teacherTuition string) (string, error) {

	if strings.TrimSpace(teacherTuition) == "" {
		return "", ErrInvalidTeacherTuition
	}

	return teacherTuition, nil
}
