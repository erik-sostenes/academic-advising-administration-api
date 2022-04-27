package schedule

import (
	"context"
	"fmt"
	"testing"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository"
)

var testScheduleRepositoryCreateSchedule = struct{
	name            string
	scheduleId      string
	scheduleAt      string
	fromDate        string
	toDate          string
	teacherTuititon string
	errExpect       error
}{
		name: "Test 1. StatusBadRequest: Invalid Fields Error",
		scheduleId: "ea58a0bc-c3a6-11ec-9d64-0242ac120002",
		scheduleAt: "2006-01-20 15:04:05",
		fromDate: "2006-01-02 15:04:05",
		toDate: "2006-01-02 15:04:05",
		teacherTuititon: "APM7392HH",
}

func Test_ScheduleRepository_CreateSchedule(t *testing.T) {
	tt := testScheduleRepositoryCreateSchedule
		
	t.Run(tt.name, func(t *testing.T) {
		schedule, gotErr := model.NewMockSchedule(tt.scheduleId, tt.scheduleAt, tt.fromDate, tt.toDate, tt.teacherTuititon, 0)
	
		assertNotError(t, nil, gotErr)

		DB := repository.NewDB()
		rep := NewScheduleStorer(DB)
		errExpect  := rep.StoreCreateSchedule(context.Background(), schedule)
			
		if _, ok := errExpect.(model.StatusBadRequest); !ok {
			t.Fatalf("expected an error of type %T, got an error of type %T", model.StatusBadRequest("") ,errExpect)
		}
	})
}

var testScheduleRepositoryStoreGetSchedulesByTeacherTuition = map[string]struct{
	teacherTuititon string
	isActive        bool
	expectedError   error
	expectedType    model.MockTeacherSchedules
}{
	"Test 1. Get a listing of model.MockTeacherSchedules{} correctly.": {
		teacherTuititon: "JFDSKR342",
		isActive: true,
		expectedError: nil,
		expectedType: model.MockTeacherSchedules{},
	},
	"Test 2. Get a listing of model.MockTeacherSchedules{} correctly.": {
		teacherTuititon: "JDK334222L",
		isActive: false,
		expectedError: nil,
		expectedType: model.MockTeacherSchedules{},
	},
} 

func Test_ScheduleRepository_StoreGetSchedulesByTeacherTuition(t *testing.T) {
	for name, tt := range testScheduleRepositoryStoreGetSchedulesByTeacherTuition {
		
		t.Run(name, func(t *testing.T) {
			DB := repository.NewDB()
			rep := NewScheduleStorer(DB)
			gotType, gotError := rep.StoreGetSchedulesByTeacherTuition(context.Background(), "APM73", true)
		
			assertNotError(t, tt.expectedError, gotError)
			
			asserType(t, tt.expectedType, gotType)
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
func asserType(t testing.TB, expetedType, gotType interface{}) {
	t.Helper()

	if !(fmt.Sprintf("%T", expetedType) == fmt.Sprintf("%T", gotType)) {
		t.Fatalf("expected structure of type %T, got structure of type %T", expetedType, gotType)
	}
}
