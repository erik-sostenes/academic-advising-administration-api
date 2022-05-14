package teacher

import (
	"context"
	"strings"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/teacher"
)
// TeacherService contains the methods that are responsible for verifying that the business logic is correct 
type TeacherService interface {
	// FindTeachers method that returns a collection of teachers with the requirements to be requested
	FindTeachers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error)
}
// teacherService implements the TeacherService interface
type teacherService struct {
	teacherStorer teacher.TeacherStorer
}
// NewTeacherService returns the default TeacherService interface implementation
func NewTeacherService(teacherStorer teacher.TeacherStorer) TeacherService {
	return &teacherService{
		teacherStorer: teacherStorer,
	}
}

func (s teacherService) FindTeachers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error) {
	if err := s.checkQueryParameters(subjectId, universityCourseId); err != nil{
		return model.TeacherCards{}, err
	}

	return s.teacherStorer.StorageFindTechers(ctx, subjectId, universityCourseId)
}

// checkQueryParameters method that verifies the query parameters are correct
// if they are not correct, an error of type StatusBadRequest is returned
func (s teacherService) checkQueryParameters(subjectId, universityCourseId string) (error) {
	if strings.TrimSpace(subjectId) == "" {
		return model.StatusBadRequest("Please verify that the value of the ´subject_id' field is correct.")
	}

	if strings.TrimSpace(universityCourseId) == "" {
		return model.StatusBadRequest("Please verify that the value of the ´university_course_id' field is correct.")
	}
	return nil
}
