package teacher

import (
	"context"
	"fmt"
	"testing"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository"
)
// TODO: add comments and refactor test unit
var testTeacherStorer = map[string]struct{
	configuration      repository.Configuration
	subjectId          string
	universityCourseId string
	expectedType       model.TeacherCards
	expectedError      error
} {
	"find successful teachers 'MySQL'": {
		configuration: repository.Configuration{
			Type: repository.SQL,
			SQL: repository.NewMySQL(),
		},
		subjectId: "JFDSKR342",
		universityCourseId: "EWJHGSH",
		expectedError: nil,
	},
}

func Test_TeacherStorer_Find(t *testing.T) {
	for name, tt := range testTeacherStorer {
		tt := tt
		t.Run(name, func(t *testing.T) {
			teacherStorer := NewSqlTeacherStorer(tt.configuration)

			got, err := teacherStorer.Find(context.TODO(), tt.subjectId, tt.universityCourseId)
			if err != tt.expectedError {
				t.Fatalf("expected error %v, got error %v", tt.expectedError, err)
			}

			if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.expectedType) {
				t.Fatalf("expected %T, got %T", tt.expectedType, got)
			}
		})
	}
} 

var testCacheTeacherStorer = []struct{
	name               string
	configuration      repository.Configuration
	subjectId          string
	universityCourseId string
	expectedType       model.TeacherCards
	expectedError      error
} {
	{
		name: "find and save successful teachers 'Redis'",
		configuration: repository.Configuration{
			Type: repository.NoSQL,
			NoSQL: repository.NewRedis(),
		},
		subjectId: "JFDSKR342",
		universityCourseId: "EWJHGSH",
		expectedError: nil,
	},
}

func Test_CacheTeacherStorer(t *testing.T) {
	for _, tt := range testCacheTeacherStorer {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cacheTeacherStorer := NewCacheTeacherStorer(tt.configuration)
			
			data := model.TeacherCard{
				Tuition: "JHJH",
				Name: "KJHKJ",
				Surnames: "JKHKJ",
			}
			err := cacheTeacherStorer.Save(context.TODO(), tt.subjectId, tt.universityCourseId, 
			model.TeacherCards{data, data},
		)
			if err != nil{
				t.Fatal(err)
			}	
			got, err := cacheTeacherStorer.Find(context.TODO(), tt.subjectId, tt.universityCourseId)
			if err != tt.expectedError {
				t.Fatalf("expected error %v, got error %v", tt.expectedError, err)
			}

			if fmt.Sprintf("%T", got) != fmt.Sprintf("%T", tt.expectedType) {
				t.Fatalf("expected %T, got %T", tt.expectedType, got)
			}
		})
	}
}
