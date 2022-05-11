package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// UserStorer interface containing the methods to interact with the MySQL database
type UserStorer interface {
	// StorageGetStudentPasswordByTuition method to get a student password
	StorageGetStudentPaswordByTuition(ctx context.Context, tuition string) (string, error)
	// StorageGetTeacherPasswordByTuition method to get a teacher password
	StorageGetTeacherPasswordByTuition(ctx context.Context, tuition string) (string, error) 
}

type userStorer struct {
	DB *sql.DB
}

// NewUserStorer returns a implements the UserStorer interface
func NewUserStorer(DB *sql.DB) UserStorer {
	return &userStorer{
		DB: DB,
	}
}

func (u *userStorer) StorageGetStudentPaswordByTuition(ctx context.Context, tuition string) (string, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	
	var password string
	if err := u.DB.QueryRowContext(queryCtx, selectStudentByTuition, tuition).Scan(
		&password,
	); err != nil {
		if err == sql.ErrNoRows {
			return "", model.NotFound(fmt.Sprintf("Student with tuition %s not found", tuition))
  	}
  	return "", model.InternalServerError(err.Error())
	}

	return password, nil 
}

func (u *userStorer) StorageGetTeacherPasswordByTuition(ctx context.Context, tuition string) (string, error) {
	queryCtx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	
	var password string
	if err := u.DB.QueryRowContext(queryCtx, selectTeacherByTuition, tuition).Scan(
		&password,
	); err != nil {
		if err == sql.ErrNoRows {
			return "", model.NotFound(fmt.Sprintf("Teacher with tuition %s not found", tuition))
  	}
  	return "",  model.InternalServerError(err.Error())
	}

	return password, nil 
}
