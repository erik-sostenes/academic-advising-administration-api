package model

import (
	"regexp"
	"strings"
)

// ErrInvalidTuition error user tuition is invalid
var ErrInvalidTuition = StatusBadRequest("Invalid tuition.")
// ErrInvalidPassword error user password is invalid
var ErrInvalidPassword = StatusBadRequest("Invalid password.")
// ErrInvalidEmail error user email is invalid
var ErrInvalidEmail   = StatusBadRequest("The email domain is invalid, it must be 'itsoeh.edu.mx'.")

// validateEmail regexp for validate email
var validateEmail = regexp.MustCompile(`[a-z0-9]+@[itsoeh]+\.[.edu]+\.[.mx]`)

// MockLogin represents the mock of a user login
type MockLogin struct {
	tuition  string 
	email    string 
	password string 
}

// NewMockLogin returns an instance of MockLogin if everything is correct
func NewMockLogin(tuition, email, password string) (MockLogin, error) {
	tuition, err := setTuition(tuition);
	if err != nil {
		return MockLogin{}, err
	}

	email, err = setEmail(email)
	if err != nil {
		return MockLogin{}, err
	}
	
	password, err = setPassword(password)
	if err != nil {
		return MockLogin{}, err
	}

	return MockLogin{
	tuition: tuition,
	email: email,
	password: password,
	}, nil
}
// Tuition represents the unique identifier of user login
func (m *MockLogin) Tuition() string {
	return m.tuition
} 
// Email represents the audience claiming the token
func (m *MockLogin) Email() string {
	return m.email
}
// Password represents user access
func (m *MockLogin) Password() string {
	return m.password
}

// setTuition returns the user tuition if everything if correct
func setTuition(tuition string) (string, error) {
	if strings.TrimSpace(tuition) == "" {
		return "", ErrInvalidTuition
	}
	return tuition, nil
}
// setPassword returns the password if everything if correct
func setPassword(password string) (string, error) {
	if strings.TrimSpace(password) == "" {
		return "", ErrInvalidPassword
	}
	return password, nil
}
// newEmail validate the email that your domain is correct 'itsoeh.edu.mx'
func setEmail(email string) (string, error){
	if !validateEmail.MatchString(email) {
		return "", ErrInvalidEmail 
	}
	return email, nil
}
