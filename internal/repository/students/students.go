package students

import (
	"context"
	"database/sql"
	"time"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
) 

type StudentsStorer interface {
	// StorageFindStudentRequest method that obtains the requests of all the students who have requested to take an advisory
	StorageFindStudentRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error)
	// StorageFindStudentRequestsAccepted method that obtains from all students that their requests have been accepted
	StorageFindStudentRequestsAccepted(ctx context.Context, teacherTuition string) (model.StudentRequestsAccepted, error)
}

type studentsStorer struct {
	DB *sql.DB	
}

func NewStudetnStorer(DB *sql.DB) StudentsStorer {
	return &studentsStorer{
		DB: DB,
	}
}

func (s *studentsStorer) StorageFindStudentRequests(ctx context.Context, teacherTuition string) (model.StudentRequests, error) {
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

func (s *studentsStorer) StorageFindStudentRequestsAccepted(ctx context.Context, teacherTuition string) (model.StudentRequestsAccepted, error) {
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
