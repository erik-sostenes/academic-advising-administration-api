package teacher

import (
	"context"
	"database/sql"
	"time"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// TeacherStorer interface containing the methods to interact with the MySQL database
type TeacherStorer interface {
	// StorageFindTechers method that seeks teachers with the requirements that are needed  
	StorageFindTechers(ctx context.Context, subjectId, universityCourseId string) (model.TeacherCards, error)
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
			&teacherCard.Name,
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
