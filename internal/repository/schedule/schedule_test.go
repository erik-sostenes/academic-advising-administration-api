package schedule

import (
	"context"
	"errors"
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
		name: "",
		scheduleId: "ea58a0be-c3a6-11ec-9d64-0242ac120002",
		scheduleAt: "2006-01-02 15:04:05",
		fromDate: "2006-01-02 15:04:05",
		toDate: "2006-01-02 15:04:05",
		teacherTuititon: "APM7392HH",
		errExpect: model.StatusBadRequest("Check that all information fields of the advisory are correct."),
}

func Test_ScheduleRepository_CreateSchedule(t *testing.T) {
	tt := testScheduleRepositoryCreateSchedule
		
	t.Run(tt.name, func(t *testing.T) {
		schedule, err := model.NewMockSchedule(tt.scheduleId, tt.scheduleAt, tt.fromDate, tt.toDate, tt.teacherTuititon)
		if err != nil {
			t.Error(err)
		}

		DB := repository.NewDB()
		rep := NewScheduleStorer(DB)
		err  = rep.StoreCreateSchedule(context.Background(), schedule)
			
		if !errors.Is(err, tt.errExpect) {
			t.Fatal(err)
		}
	})
}
