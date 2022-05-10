package user

import (
	"context"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/user"
	"github.com/itsoeh/academy-advising-administration-api/internal/services"
)
// ScheduleService contains the methods that are responsible for verifying that the business logic is correct
// and give authorization to the user, if it meets all the requirements
type UserService interface {
	// GetStudentPasswordByTuition method that checks if the student meets all the requirements
	GetStudentPasswordByTuition(ctx context.Context, tuition, email, password string) (string, error)
	// GetTeacherPasswordByTuition method that verifying if it teacher meets all the requirements
	GetTeacherPasswordByTuition(ctx context.Context, tuition, email, password string) (string, error)
}
// userService implements the UserService interface
type userService struct {
	userStorer user.UserStorer
	encrytor services.Encrytor
	token services.Token
}
// NewUserService returns the default UserService interface implementation
func NewUserService(userStorer user.UserStorer) UserService {
	return &userService{
		userStorer: userStorer,
		encrytor: services.Encrytor{},
		token: services.Token{},
	}
}

func (s *userService) GetStudentPasswordByTuition(ctx context.Context, tuition, email, password string) (tokenString string, err error) {
	mockLogin, err := model.NewMockLogin(tuition, email, password)	
	if err != nil {
		return 
	}
	
	password, err = s.userStorer.StorageGetStudentPaswordByTuition(ctx, mockLogin.Tuition())
	if err != nil {
		return
	}
	
	if err = s.encrytor.ValidatePassword(password, mockLogin.Password()); err != nil {
		return
	}
	
	tokenString, err = s.token.GeterateToken(mockLogin.Email())	
	if err != nil {
		return 
	}

	return 
}

func (s *userService) GetTeacherPasswordByTuition(ctx context.Context, tuition, email, password string) (tokenString string, err error) {
		mockLogin, err := model.NewMockLogin(tuition, email, password)	
	if err != nil {
		return 
	}
	
	password, err = s.userStorer.StorageGetTeacherPasswordByTuition(ctx, mockLogin.Tuition())
	if err != nil {
		return
	}
	
	if err = s.encrytor.ValidatePassword(password, mockLogin.Password()); err != nil {
		return
	}
	
	tokenString, err = s.token.GeterateToken(mockLogin.Email())	
	if err != nil {
		return 
	}

	return 
}



