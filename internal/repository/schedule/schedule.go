package schedule

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// ScheduleStorer interface containing the methods to interact with the MySQL database
type ScheduleStorer interface {
	// StoreCreateSchedule method that create a new teacher schedule 
	StorageCreateSchedule(context.Context, model.MockSchedule) error
	// StoreGetSchedulesByTeacherTuition method that gets a collection of teacher schedules by teacherId and if the schedule is active. 
	StorageGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive bool) (model.TeacherSchedules, error) 
}

type scheduleStorer struct {
	DB *sql.DB
}

// NewScheduleStorer returns a structure that implements the ScheduleStorer interface
func NewScheduleStorer(DB *sql.DB) ScheduleStorer {
	return &scheduleStorer{
		DB: DB,
	}
}

func (s *scheduleStorer) StorageCreateSchedule(ctx context.Context, schedule model.MockSchedule) (err error) {
	queryCTx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel() 

	_, err = s.DB.ExecContext(queryCTx,insertTeacherSchedules, 
		schedule.ScheduleId(),
		schedule.ScheduleAt(),
		schedule.FromDate(),
		schedule.ToDate(),
		schedule.TeacherTuition(),
	)
	
	if code, ok := err.(*mysql.MySQLError); ok {
		//NOTE: Error Code: 1062. Duplicate entry 'value' for key 'schedule_at'
		//NOTE: Error Code: 1452. Cannot add or update a child row: a foreign key constraint fails

		if code.Number == 1062 {
			err = model.StatusBadRequest(fmt.Sprintf("A schedule whith date %v was already found.", schedule.ScheduleAt()))
			return
		}
		
		if code.Number == 1452 {
			err = model.StatusBadRequest("Check that all information fields of the schedule are correct.")
			return
		}
		
		err = model.InternalServerError("An error has occurred when adding a new schedule.")
		return
	}

	return
}

func (s *scheduleStorer) 	StorageGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive bool) ( model.TeacherSchedules, error) {
	queryCTx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	
	rows, err := s.DB.QueryContext(queryCTx, selectTeacherSchedulesByTeacherTuition, teacherId, isActive)
	defer rows.Close()

	if err != nil {
		err = model.InternalServerError("An error has ocurred while obtainig the teacher schedule.")
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
			err = model.InternalServerError(fmt.Sprintf("Teacher schedule: %v", err))
			return model.TeacherSchedules{}, err
		}

		teacherSchedules = append(teacherSchedules, teacherSchedule)
	}
	return teacherSchedules, err	
} 

