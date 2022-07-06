package student

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

func TestStudentStorer_Find(t *testing.T) {
	SQL := repository.NewMySQL()
		
	studentStorer := NewStudetnStorer(repository.Configuration{
		SQL: SQL,
		Type: repository.SQL,
	})

	t.Cleanup(func() {	
		SQL.Close()	
	})
	
	tsc := map[string] struct{
		teacherTuition          				string
		expectedStudentRequests 				model.StudentRequests
		expectedStudentAcceptedRequests model.StudentAcceptedRequests
		expectedError                   error
	}{
		"Test 1: find successful student requests 'MySQL'": {
			teacherTuition: "some_teacher_tuition",
			expectedError: nil,
		},
		"Test 2: find successful student requests 'MySQL'": {
			teacherTuition: "some_teacher_tuition",
			expectedError: nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			gotRequests, gotErr := studentStorer.FindRequests(context.TODO(), ts.teacherTuition)
			assertNotError(t, ts.expectedError, gotErr)

			asserType(t, ts.expectedStudentRequests, gotRequests)

			gotAcceptedRequests, gotErr := studentStorer.FindAcceptedRequests(context.TODO(), ts.teacherTuition)
			assertNotError(t, ts.expectedError, gotErr)

			asserType(t, ts.expectedStudentAcceptedRequests, gotAcceptedRequests)
		})
	}
}

func TestCacheStudentStorer_Find(t *testing.T) {
	NoSQL := repository.NewRedis()
		
	studentStorer := NewCacheStudentStorer(repository.Configuration{
		NoSQL: NoSQL,
		Type: repository.NoSQL,
	})

	tsc := map[string] struct{
		teacherTuition          				string
		expectedStudentRequests 				model.StudentRequests
		expectedStudentAcceptedRequests model.StudentAcceptedRequests
		expectedError                   error
	}{
		"Test 1: find successful student requests 'Redis'": {
			teacherTuition: "some_teacher_tuition",
			expectedError: redis.Nil,
		},
		"Test 2: find successful student requests 'Redis'": {
			teacherTuition: "some_teacher_tuition",
			expectedError: redis.Nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			gotRequests, gotErr := studentStorer.FindRequests(context.TODO(), ts.teacherTuition)
			assertNotError(t, ts.expectedError, gotErr)

			asserType(t, ts.expectedStudentRequests, gotRequests)

			gotAcceptedRequests, gotErr := studentStorer.FindAcceptedRequests(context.TODO(), ts.teacherTuition)
			assertNotError(t, ts.expectedError, gotErr)

			asserType(t, ts.expectedStudentAcceptedRequests, gotAcceptedRequests)
		})
	}
}

func TestCacheStudentStorer_Save(t *testing.T) {
	NoSQL := repository.NewRedis()
		
	studentStorer := NewCacheStudentStorer(repository.Configuration{
		NoSQL: NoSQL,
		Type: repository.NoSQL,
	})

	tsc := map[string] struct {
		teacherTuition          string
		studentRequests 				model.StudentRequests
		studentAcceptedRequests model.StudentAcceptedRequests
		expectedError           error
	}{
		"Test 1: save successful student requests 'Redis'": {
			teacherTuition: "some_teacher_tuition",
			studentRequests: model.StudentRequests{
				model.StudentRequest {
					Tuition:           "some_tuition_1",
					Name:              "some_name_1",
					Email:             "some_email_1",
					CubicleNumber:     1,
					Subject:           "some_subject_1",
					BuildingNumber:    "some_building_number_1",
					AdvisoryId:        "some_advisory_id_1",
					TeacherScheduleId: "some_teacher_schedule_id_1",
				},
				model.StudentRequest {
					Tuition:           "some_tuition_2",
					Name:              "some_name_2",
					Email:             "email",
					CubicleNumber:     2,
					Subject:           "some_subject_2",
					BuildingNumber:    "some_building_number_2",
					AdvisoryId:        "some_advisory_id_2",
					TeacherScheduleId: "some_teacher_schedule_id_2",
				},
			},
			studentAcceptedRequests: model.StudentAcceptedRequests{
				model.StudentAcceptedRequest{
					Tuition:          "some_tuition_1",
					Name:             "some_name_1",
					Email:            "some_email_1",
					CubicleNumber:    1,
					Subject:          "some_subject_1",
					UniversityCourse: "some_university_course_1",
				},
				model.StudentAcceptedRequest{
					Tuition:          "some_tuition_2",
					Name:             "some_name_2",
					Email:            "some_email_2",
					CubicleNumber:    1,
					Subject:          "some_subject_2",
					UniversityCourse: "some_university_course_2",
				},
			},
			expectedError: nil,
		},
		"Test 2: save successful student requests 'Redis'": {
			teacherTuition: "some_teacher_tuition",
			studentRequests: model.StudentRequests{
				model.StudentRequest {
					Tuition:           "some_tuition_1",
					Name:              "some_name_1",
					Email:             "some_email_1",
					CubicleNumber:     1,
					Subject:           "some_subject_1",
					BuildingNumber:    "some_building_number_1",
					AdvisoryId:        "some_advisory_id_1",
					TeacherScheduleId: "some_teacher_schedule_id_1",
				},
				model.StudentRequest {
					Tuition:           "some_tuition_2",
					Name:              "some_name_2",
					Email:             "email",
					CubicleNumber:     2,
					Subject:           "some_subject_2",
					BuildingNumber:    "some_building_number_2",
					AdvisoryId:        "some_advisory_id_2",
					TeacherScheduleId: "some_teacher_schedule_id_2",
				},
			},
			studentAcceptedRequests: model.StudentAcceptedRequests{
				model.StudentAcceptedRequest{
					Tuition:          "some_tuition_1",
					Name:             "some_name_1",
					Email:            "some_email_1",
					CubicleNumber:    1,
					Subject:          "some_subject_1",
					UniversityCourse: "some_university_course_1",
				},
				model.StudentAcceptedRequest{
					Tuition:          "some_tuition_2",
					Name:             "some_name_2",
					Email:            "some_email_2",
					CubicleNumber:    1,
					Subject:          "some_subject_2",
					UniversityCourse: "some_university_course_2",
				},
			},
			expectedError: nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			gotErr := studentStorer.SaveRequests(context.TODO(), ts.teacherTuition, ts.studentRequests)
			assertNotError(t, ts.expectedError, gotErr)

			gotErr = studentStorer.SaveAcceptedRequests(context.TODO(), ts.teacherTuition, ts.studentAcceptedRequests)
			assertNotError(t, ts.expectedError, gotErr)
		})
	}
}

// assertNotError asserts that the error will be nil
 func assertNotError(t testing.TB, expectedError, gotError error) {
	t.Helper()

	if expectedError != gotError {
		t.Fatalf("expected error %v, got an error %v", expectedError, gotError)
	}
}

// asserType asserts that the object is of the requested type
func asserType(t testing.TB, expectedType, gotType interface{}) {
	t.Helper()

	if  reflect.TypeOf(expectedType) != reflect.TypeOf(gotType) {
		t.Fatalf("expected %T, got %T", expectedType, gotType)
	}
}
