package schedule

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
)

// ScheduleStorer interface containing the methods to interact with the MySQL database
type ScheduleStorer interface {
	// StoreCreateSchedule method that create a new teacher schedule 
	StoreCreateSchedule(context.Context, model.MockSchedule) error
	// StoreGetSchedulesByTeacherTuition method that gets a collection of teacher schedules by teacherId and if the schedule is active. 
	StoreGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive bool) (model.MockTeacherSchedules, error) 
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

func (s *scheduleStorer) StoreCreateSchedule(ctx context.Context, schedule model.MockSchedule) (err error) {
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

func (s *scheduleStorer) 	StoreGetSchedulesByTeacherTuition(ctx context.Context, teacherId string, isActive bool) (mock model.MockTeacherSchedules, err error) {
	queryCTx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	
	rows, err := s.DB.QueryContext(queryCTx, selectTeacherSchedulesByTeacherTuition, teacherId, isActive)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		err = model.InternalServerError("An error has ocurred while obtainig the teacher schedule.")
		return 
	}

	for rows.Next() {
		var day, teacherName, surnameTeacher string
		var scheduleId, scheduleAt, fromDate, toDate, teacherTuition string	
		var studentAccountant uint8

		if err = rows.Scan(
			&day,
			&teacherName,
			&surnameTeacher,
			&scheduleId,
			&scheduleAt,
			&fromDate,
			&toDate,
			&studentAccountant,
			&teacherTuition,
		); err != nil {
			err = model.InternalServerError(fmt.Sprintf("teacher schedule: %v", err))
			return
		}

		mockSchedule, errM  := model.NewMockSchedule(scheduleId, scheduleAt, fromDate, toDate, teacherTuition, studentAccountant)
		if errM != nil {
			err = errM
			return 
		}
		
		mockTeacherSchedule, errM := model.NewMockTeacherSchedule(day, teacherName, surnameTeacher, mockSchedule)
		if errM != nil {
			err = errM
			return 
		}

		mock = append(mock, mockTeacherSchedule)
	}
	return	
} 

