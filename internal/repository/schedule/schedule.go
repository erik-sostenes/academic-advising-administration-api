package schedule

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

// ScheduleStorer interface containing the methods to interact with the MySQL database
type ScheduleStorer interface {
	// Save method that create a new teacher schedule 
	Save(context.Context, model.MockSchedule) error
	// Find method that gets a collection of teacher schedules by teacherId and if the schedule is active. 
	Find(ctx context.Context, teacherId string, isActive bool) (model.TeacherSchedules, error) 
}

type scheduleStorer struct {
	DB *sql.DB
}

// NewScheduleStorer returns a structure that implements the ScheduleStorer interface
func NewScheduleStorer(c repository.Configuration) ScheduleStorer {
	switch c.Type {
	case repository.SQL:
		return &scheduleStorer{
			DB: c.SQL,
		}
	default:
		panic(fmt.Sprintf("%T type is not supported", repository.SQL))
	} 
}

func (s *scheduleStorer) Save(ctx context.Context, schedule model.MockSchedule) (err error) {
	queryCTx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel() 

	_, err = s.DB.ExecContext(queryCTx,insertTeacherSchedules, 
		schedule.ScheduleId(),
		schedule.ScheduleAt(),
		schedule.FromDate(),
		schedule.ToDate(),
		schedule.TeacherTuition(),
	)
	return
}

func (s *scheduleStorer) Find(ctx context.Context, teacherId string, isActive bool) (model.TeacherSchedules, error) {
	queryCTx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	
	rows, err := s.DB.QueryContext(queryCTx, selectTeacherSchedulesByTeacherTuition, teacherId, isActive)
	defer rows.Close()

	if err != nil {
		return model.TeacherSchedules{}, err
	}
	
	var teacherSchedules model.TeacherSchedules

	for rows.Next() {
		var teacherSchedule model.TeacherSchedule

		if err = rows.Scan(
			&teacherSchedule.Day,
			&teacherSchedule.TeacherName,
			&teacherSchedule.SurnameTeacher,
			&teacherSchedule.Schedule.ScheduleId,
			&teacherSchedule.Schedule.ScheduleAt,
			&teacherSchedule.Schedule.FromDate,
			&teacherSchedule.Schedule.ToDate,
			&teacherSchedule.Schedule.StudentAccountant,
			&teacherSchedule.Schedule.TeacherTuition,
		); err != nil {
			return model.TeacherSchedules{}, err
		}
		teacherSchedules = append(teacherSchedules, teacherSchedule)
	}
	return teacherSchedules, err	
} 
