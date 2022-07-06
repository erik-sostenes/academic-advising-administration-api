package teacher

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)

func Test_TeacherStorer_Find(t *testing.T) {
	var testTeacherStorer = map[string]struct{
		teacherStorer      TeacherStorer 
		subjectId          string
		universityCourseId string
		expectedType       model.TeacherCards
		expectedError      error
	} {
		"find successful teachers 'MySQL'": {
			teacherStorer: NewSqlTeacherStorer(
				repository.Configuration{
					Type: repository.SQL,
					SQL: repository.NewMySQL(),
				},
			),
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: nil,
		},
		"find successful teachers 'Redis'": {
			teacherStorer: NewCacheTeacherStorer(
				repository.Configuration{
					Type: repository.NoSQL,
					NoSQL: repository.NewRedis(),
				},
			),
			subjectId: "JFDSKR342",
			universityCourseId: "EWJHGSH",
			expectedError: redis.Nil,
		},	
	}
	for name, tt := range testTeacherStorer {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got, err := tt.teacherStorer.Find(context.TODO(), tt.subjectId, tt.universityCourseId)
			if err != tt.expectedError {
				t.Fatalf("expected error %v, got error %v", tt.expectedError, err)
			}

			if reflect.TypeOf(got) != reflect.TypeOf(tt.expectedType) {
				t.Fatalf("expected %T, got %T", tt.expectedType, got)
			}
		})
	}
} 

func Test_CacheTeacherStorer_Save(t *testing.T) {
	var testCacheTeacherStorer = map[string]struct {
		cacheTeacherStorer CacheTeacherStorer
		subjectId,
		universityCourseId string
		teacherCards       model.TeacherCards
		expectedError      error
	} {
		"save successful teachers 'Redis' 1" : {
			cacheTeacherStorer: NewCacheTeacherStorer (
				repository.Configuration {
					Type: repository.NoSQL,
					NoSQL: repository.NewRedis(),
				},
			),
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
			cacheTeacherStorer: NewCacheTeacherStorer (
				repository.Configuration {
					Type: repository.NoSQL,
					NoSQL: repository.NewRedis(),
				},
			),
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

	for name, tt := range testCacheTeacherStorer {
		t.Run(name, func(t *testing.T) {
			err := tt.cacheTeacherStorer.Save(context.TODO(),
				tt.subjectId,
				tt.universityCourseId,
				tt.teacherCards,
			)
			if err != tt.expectedError {
				t.Fatalf("expected error %v, got error %v", tt.expectedError, err)
			}
		})
	}
}
