package schedule

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

func TestScheduleStorer_Save(t *testing.T) {
	var tsc = map[string]struct{
		schedule            model.Schedule
		expectedError       error
		expectedMySQLError  *mysql.MySQLError
	}{
		"Test 1: Save 'MySQL' schedule successful": {
			schedule: model.Schedule {
				ScheduleId: "ea58a0bc-c3a6-11ec-9d64-0242ac120002",
				ScheduleAt: "2006-01-20 15:04:05",
				FromDate: "2006-01-02 15:04:05",
				ToDate: "2006-01-02 15:04:05",
				StudentAccountant: 3,
				TeacherTuition: "APM7392HH",
			},
			expectedError: nil,
		},
		"Test 2: Save 'MySQL' schedule successful": {
			schedule: model.Schedule {
				ScheduleId: "ea59a0bc-c3a6-11ec-9d64-0242ac120002",
				ScheduleAt: "2006-01-20 15:04:05",
				FromDate: "2006-01-02 15:04:05",
				ToDate: "2006-01-02 15:04:05",
				StudentAccountant: 34,
				TeacherTuition: "APM7392HH",
			},
			expectedError: nil,
		},
	}		

	SQL := repository.NewMySQL()

	scheduleStorer := NewScheduleStorer(repository.Configuration{
		SQL: SQL,
		Type: repository.SQL,
	})
	
	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			schedule, gotErr := model.NewMockSchedule(
				ts.schedule.ScheduleId,
				ts.schedule.ScheduleAt,
				ts.schedule.FromDate,
				ts.schedule.ToDate,
				ts.schedule.TeacherTuition,
				ts.schedule.StudentAccountant,
			)	
			assertNotError(t, ts.expectedError, gotErr)

			gotError := scheduleStorer.Save(context.Background(), schedule)
			
			_, ok := gotError.(*mysql.MySQLError)
			if !ok {
				t.Fatalf("expected %v type error, got %v type error", ts.expectedMySQLError, gotError)
			}
		})
	}
}

func TestScheduleStorer_Find(t *testing.T) {
	var tsc = map[string]struct{
		teacherId       string
		isActive        bool
		expectedError   error
		expectedType    model.TeacherSchedules
	}{
		"Test 1: save successful schedule 'MySQL'": {
			teacherId: "JFDSKR342",
			isActive: true,
			expectedError: nil,
		},
		"Test 2: save successful schedule 'MySQL'": {
			teacherId: "JDK334222L",
			isActive: false,
			expectedError: nil,
		},
	} 
	
	SQL := repository.NewMySQL()
	scheduleStorer := NewScheduleStorer(repository.Configuration{
		SQL: SQL,
		Type: repository.SQL,
	})
	
	for name, ts := range tsc {	
		t.Run(name, func(t *testing.T) {
			gotType, gotError := scheduleStorer.Find(context.Background(), ts.teacherId, ts.isActive)
		
			assertNotError(t, ts.expectedError, gotError)
			
			assertType(t, ts.expectedType, gotType)
		})
	}
}

// assertNotError asserts that the error will be null
 func assertNotError(t testing.TB, expected, got error) {
	t.Helper()

	if expected != got {
		t.Fatalf("expected error %v, got an error %v", expected, got)
	}
}

// asserType asserts that the object is of the requested type
func assertType(t testing.TB, expectedType, gotType interface{}) {
	t.Helper()

	if  reflect.TypeOf(expectedType) != reflect.TypeOf(gotType) {
		t.Fatalf("expected %T, got %T", expectedType, gotType)
	}
}
