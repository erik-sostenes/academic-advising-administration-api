package schedule

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
)

// ScheduleStorer interface that contains the method to store a new teacher schedule to the database
type ScheduleStorer interface {
	// StoreCreateSchedule method that create a new teacher schedule 
	StoreCreateSchedule(context.Context, model.MockSchedule) error
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
			err = model.StatusBadRequest("Check that all information fields of the advisory are correct.")
			return
		}
		
		err = model.InternalServerError("An error has occurred when adding a new teacher schedule.")
		return
	}

	return
}
