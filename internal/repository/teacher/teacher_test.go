package teacher

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

func TestTeacherStorer_Find(t *testing.T) {
	SQL := repository.NewMySQL()

	teacherStorer := NewSqlTeacherStorer(
		repository.Configuration {
			Type: repository.SQL,
			SQL: SQL,
		},
	)
	
	t.Cleanup(func() {
		SQL.Close()
	})

	var tsc = map[string]struct{
		subjectId          string
		universityCourseId string
		expectedType       model.TeacherCards
		expectedError      error
	} {
		"find successful teachers 'MySQL'": {
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: nil,
		},
		"find successful teachers 'Redis'": {
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: nil,
		},	
	}

	for name, ts := range tsc {
		ts := ts
		t.Run(name, func(t *testing.T) {
			gotType, gotErr := teacherStorer.Find(context.TODO(), ts.subjectId, ts.universityCourseId)

			assertNotError(t, ts.expectedError, gotErr)
			
			assertType(t, ts.expectedType, gotType)
		})
	}
} 

func TestCacheStudentStorer_Find(t *testing.T) {
	NoSQL := repository.NewRedis()

	teacherStorer := NewCacheTeacherStorer(
		repository.Configuration {
			Type: repository.NoSQL,
			NoSQL: NoSQL,
		},
	)

	var tsc = map[string]struct{
		subjectId          string
		universityCourseId string
		expectedType       model.TeacherCards
		expectedError      error
	} {
		"find successful teachers 'MySQL'": {
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: redis.Nil,
		},
		"find successful teachers 'Redis'": {
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: redis.Nil,
		},	
	}

	for name, ts := range tsc {
		ts := ts
		t.Run(name, func(t *testing.T) {
			gotType, gotErr := teacherStorer.Find(context.TODO(), ts.subjectId, ts.universityCourseId)

			assertNotError(t, ts.expectedError, gotErr)
			
			assertType(t, ts.expectedType, gotType)
		})
	}
} 

func TestCacheTeacherStorer_Save(t *testing.T) {
	NoSQL := repository.NewRedis()

	cacheTeacherStorer := NewCacheTeacherStorer(
		repository.Configuration {
			Type: repository.NoSQL,
			NoSQL: NoSQL,
		},
	)

	var tsc = map[string]struct {
		subjectId,
		universityCourseId string
		teacherCards       model.TeacherCards
		expectedError      error
	} {
		"save successful teachers 'Redis' 1" : {
			subjectId: "3J3N4",
			universityCourseId: "7O3H",
			teacherCards: []model.TeacherCard{
				{
					Tuition: "12JND3",
					Name: "Carlos",
					Surnames: "Sanchez",
					Email: "carlos@gmail.com",
					IsActive: true,
					Cubicle: "3J3N4",
					SubjectId: "3J3N4",
					UniversityCourseId: "7O3H",
					SubcoordinatorTuition: "22J2H3",
					CoordinatorTuition: "11HHH3",
				},
				{
					Tuition: "33KJD",
					Name: "Jose Antonio",
					Surnames: "Mendoza",
					Email: "jose_antonio@gmail.com",
					IsActive: true,
					Cubicle: "3J389",
					SubjectId: "3J3N4",
					UniversityCourseId: "7O3H",
					SubcoordinatorTuition: "JJEI3",
					CoordinatorTuition: "97JJH",
				},
			},	
			expectedError: nil,
		},
		"save successful teachers 'Redis' 2" : {
			subjectId: "H3HJ",
			universityCourseId: "5H3H",
			teacherCards: []model.TeacherCard{
				{
					Tuition: "2345J",
					Name: "Axel Yael",
					Surnames: "Castro",
					Email: "axel_castro@gmail.com",
					IsActive: true,
					Cubicle: "UUNG",
					SubjectId: "H3HJ",
					UniversityCourseId: "5H3H",
					SubcoordinatorTuition: "732H",
					CoordinatorTuition: "8MND",
				}, 
				{
					Tuition: "797J8",
					Name: "Jonathan",
					Surnames: "Hernandez",
					Email: "jonathan@gmail.com",
					IsActive: true,
					Cubicle: "7JD8",
					SubjectId: "H3HJ",
					UniversityCourseId: "5H3H",
					SubcoordinatorTuition: "666H",
					CoordinatorTuition: "97J00",
				},
			},	
			expectedError: nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			gotErr := cacheTeacherStorer.Save(context.TODO(),
				ts.subjectId,
				ts.universityCourseId,
				ts.teacherCards,
			)
			assertNotError(t, ts.expectedError, gotErr)

			t.Cleanup(func() {
				cacheTeacherStorer.Delete(context.TODO(), ts.subjectId, ts.universityCourseId)	
			})
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

// assertType asserts that the object is of the requested type
func assertType(t testing.TB, expectedType, gotType interface{}) {
	t.Helper()

	if  reflect.TypeOf(expectedType) != reflect.TypeOf(gotType) {
		t.Fatalf("expected %T, got %T", expectedType, gotType)
	}
}
