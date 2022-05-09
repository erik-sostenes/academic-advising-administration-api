package services

import (
	"regexp"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// validateEmail regexp for validate email
var validateEmail = regexp.MustCompile(`[a-z0-9]+@[itsoeh]+\.[.edu]+\.[.mx]`)

// Validator contains all the methods for the bussnies logic
type Validator struct{}

//  ValidateEmail validate the email that your domain is correct 'itsoeh.edu.mx'
func (*Validator) ValidateEmail(email string) (err error){
	if !validateEmail.MatchString(email) {
		err = model.StatusBadRequest("the email domain is invalid, it must be 'itsoeh.edu.mx'.")
		return
	}
	return
}

// ValidateThePassword validate the encrypted password if correct
func (*Validator) ValidateThePassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err = model.Forbidden("Your password does not exist or is incorrect.")
		return
	}
	return
}
