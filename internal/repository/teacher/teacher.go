package teacher

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// TeacherStorer interface containing the methods to interact with the MySQL database
type TeacherStorer interface {
	// StorageFindTechers method that seeks teachers with the requirements that are needed  
	StorageFindTechers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error)
	// StorageFindStudentRequest method that obtains the requests of all the students who have requested to take an advisory
	StorageFindStudentRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error)
	// StorageFindStudentRequestsAccepted
	StorageFindStudentRequestsAccepted(ctx context.Context, teacherTuition string) (model.StudentRequestsAccepted, error)
}

// teacherStorer  implements TeacherStorer interface
type teacherStorer struct {
	DB *sql.DB
}

// NewTeacherStorer returns a structure that implements the TeacherStorer interface
func NewTeacherStorer(DB *sql.DB) TeacherStorer {
	return &teacherStorer{
		DB: DB,
	}
} 

func (s *teacherStorer) StorageFindTechers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectTeachersByCareerAndSubject,
		subjectId, 
		universityCourseId,
	)
	defer rows.Close()
	
	if err != nil {
		return model.TeacherCards{}, model.InternalServerError("An error has ocurred while obtainig the teachers.")
	}

	var teacherCards model.TeacherCards

	for rows.Next() {
		var teacherCard model.TeacherCard

		if err := rows.Scan(
			&teacherCard.Tuition,
			&teacherCard.Name,
			&teacherCard.Surnames,
			&teacherCard.Email,
			&teacherCard.IsActive,
			&teacherCard.Cubicle,
			&teacherCard.SubjectId,
			&teacherCard.UniversityCourseId,
			&teacherCard.SubcoordinatorTuition,
			&teacherCard.CoordinatorTuition,
		); err != nil {
			return teacherCards, model.InternalServerError("Error when searching for the teacher.")
		}
		teacherCards = append(teacherCards, teacherCard)
	}

	if err := rows.Err(); err != nil {
		return teacherCards, model.InternalServerError(err.Error())
	}

	return teacherCards, err 
}

func (s *teacherStorer) StorageFindStudentRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectStudentRequest, teacherTuition)
	defer rows.Close()
	
	if err != nil {
		return model.StudentRequests{}, model.InternalServerError("An error has ocurred while obtainig the student request.")
	}

	var studentRequests model.StudentRequests

	for rows.Next() {
		var studentRequest model.StudentRequest

		if err := rows.Scan(
			&studentRequest.Tuition,
			&studentRequest.Name,
			&studentRequest.Email,
			&studentRequest.CubicleNumber,
			&studentRequest.Subject,
			&studentRequest.BuildingNumber,
			&studentRequest.AdvisoryId,
			&studentRequest.TeacherScheduleId,
		); err != nil {
			return model.StudentRequests{}, model.InternalServerError("Error when searching for the reques students.")
		}
		studentRequests = append(studentRequests, studentRequest)
	}

	if err := rows.Err(); err != nil {
		return model.StudentRequests{}, model.InternalServerError(err.Error())
	}

	return studentRequests, err 

}

func (s *teacherStorer) StorageFindStudentRequestsAccepted(ctx context.Context, teacherTuition string) (model.StudentRequestsAccepted, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	rows, err := s.DB.QueryContext(queryCtx, selectStudentRequestAccepted, teacherTuition)
	defer rows.Close()
	
	if err != nil {
		return model.StudentRequestsAccepted{}, model.InternalServerError("An error has ocurred while obtainig the student request.")
	}

	var studentRequestsAccepted model.StudentRequestsAccepted

	for rows.Next() {
		var studentRequestAccepted model.StudentRequestAccepted

		if err := rows.Scan(
			&studentRequestAccepted.Tuition,
			&studentRequestAccepted.Name,
			&studentRequestAccepted.Email,
			&studentRequestAccepted.CubicleNumber,
			&studentRequestAccepted.Subject,
			&studentRequestAccepted.UniversityCourse,
		); err != nil {
			log.Println(err)
			return model.StudentRequestsAccepted{}, model.InternalServerError("Error when searching for the reques students.")
		}
		studentRequestsAccepted = append(studentRequestsAccepted, studentRequestAccepted)
	}

	if err := rows.Err(); err != nil {
		return model.StudentRequestsAccepted{}, model.InternalServerError(err.Error())
	}

	return studentRequestsAccepted, err 
}
