package schedule

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
)

var testCreateSchedule = map[string] struct{
	scheduleId string 
	scheduleAt string
	fromDate string
	toDate string
	teacherTuition string
	expectError error
}{
	"": {
		scheduleId: "9dac8d9c-c280-11ec-9d64-0242ac120002",
		scheduleAt: "2022-04-01T08:41:50Z",
		fromDate: "2022-04-01T08:41:50Z",
		toDate: "2022-04-01T08:41:50Z",
		teacherTuition: "128DHNR80",
		expectError: errors.New(""),
	},
	"Incorrect": {
		scheduleId: "d83fcdfc-c280-11ec-9d64-0242ac120002",
		scheduleAt: "2022-04-01T08:41:50Z",
		fromDate: "2022-04-01T08:41:50Z",
		toDate: "2022-04-01T08:41:50Z",
		teacherTuition: "128DHNR80",
		expectError: errors.New(""),
	},
}

func Test_ScheduleRepository_CreateSchedule(t *testing.T) {
	for name, tt := range testCreateSchedule {
		tt := tt
		t.Run(name, func(t *testing.T) {
			schedule, err := model.NewMockSchedule(tt.scheduleId, tt.scheduleAt, tt.fromDate, tt.toDate, tt.teacherTuition)	

			if err != nil {
				t.Log(err)
			}

			db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)))
			if err != nil {
				t.Log(err)
			}
			sqlmock.ExpectedExec(
				insertTeacherSchedules).
				WithArgs(tt.scheduleId, tt.scheduleAt, tt.fromDate, tt.toDate, tt.teacherTuition).
				WillReturnResults(sqlmock.NewResult(0,1))
		
			repo := NewScheduleRepository(db)

			err = repo.Create(schedule)
			if err != nil {
				t.Log(err)
			}
		})
	} 
}
