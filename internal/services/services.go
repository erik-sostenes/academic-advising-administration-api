package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/schedule"
)

// ScheduleService contains the methods that are responsible for verifying that the business logic is correct
type ScheduleService interface {
	// CreateSchedule returns a new model.MockSchedule 
	CreateSchedule(ctx context.Context, scheduleId, scheduleAt, fromDate, toDate, teacherTuition string) error 
	// StoreGetSchedulesByTeacherTuition  returns a new model.MockTeacherSchedules
	StoreGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive string) (model.MockTeacherSchedules, error)
}

// ScheduleService implements the ScheduleService interface
type scheduleService struct {
	scheduleStorer schedule.ScheduleStorer
}

// NewScheduleService returns the default Service interface implementation
func NewScheduleService(scheduleStorer schedule.ScheduleStorer) ScheduleService {
	return &scheduleService{
		scheduleStorer: scheduleStorer,
	}
}

func (s scheduleService) CreateSchedule(ctx context.Context, scheduleId, scheduleAt, fromDate, toDate, teacherTuition string) error {
	schedule, err := model.NewMockSchedule(scheduleId, scheduleAt, fromDate, toDate, teacherTuition, 0)
	if err != nil {
		return err
	}
	
	return s.scheduleStorer.StoreCreateSchedule(ctx, schedule)
}


func (s scheduleService) StoreGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive string) (model.MockTeacherSchedules, error) {
	teacherIdVO, isActiveVO, err :=  s.checkQueryParameters(teacherId, isActive)	
	if err != nil {
		return model.MockTeacherSchedules{}, err
	}

	return s.scheduleStorer.StoreGetSchedulesByTeacherTuition(ctx, teacherIdVO, isActiveVO)
}

// checkQueryParameters method that verifies the query parameters are correct
// if they are not correct, an error of type StatusBadRequest is returned
func (s scheduleService) checkQueryParameters(teacherId string, isActive string) (string, bool, error) {
	boolValue, err := strconv.ParseBool(isActive)
	if err != nil {
		return teacherId, boolValue, model.StatusBadRequest("Please verify that the value of the ´is_active' field is correct.")
	}

	if strings.TrimSpace(teacherId) == "" {
		return "", boolValue, model.StatusBadRequest("Please verify that the value of the ´teacher_id' field is correct.")
	}

	return teacherId, boolValue, nil
}
