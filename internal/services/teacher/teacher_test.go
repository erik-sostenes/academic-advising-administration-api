package teacher

import (
	"context"
	"errors"
	"testing"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/repository/repositorymocks/teacher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTeacherService_Find(t *testing.T) {
	tsc := map[string]struct {
		subjectId,
		universityCourseId string
		teacherCards       model.TeacherCards
		expectedError      error
	} {
		"get teachers right" : {
			subjectId: "3J3N4J8",
			universityCourseId: "7O3HJ8K",
			teacherCards: []model.TeacherCard {
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
		"get teachers wrong" : {
			subjectId: "H3HJ17TH8",
			universityCourseId: "5H3HHJS1",
			teacherCards: []model.TeacherCard{},	
			expectedError: errors.New("an unexpected error occurs"),
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			teacherStorerMock := new(repositorymocks.TeacherStorer)
			
			teacherStorerMock.On("Find",
				mock.Anything,
				ts.subjectId,
				ts.universityCourseId,
			).Return(ts.teacherCards, ts.expectedError)
		
			teacherService := NewTeacherService(teacherStorerMock)
			teacherCards, err := teacherService.Find(context.Background(), ts.subjectId, ts.universityCourseId)

			t.Cleanup(func() {	
				teacherStorerMock.AssertExpectations(t)
			})

			assert.ErrorIs(t, ts.expectedError, err)
			assert.Exactly(t, ts.teacherCards, teacherCards)
		})
	}
}

