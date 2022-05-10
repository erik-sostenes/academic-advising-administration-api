package services

import (
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Validator contains all the methods for the bussnies logic
type Validator struct{}

// ValidateThePassword validate the encrypted password if correct
func (*Validator) ValidateThePassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err = model.Forbidden("Your password is incorrect.")
		return
	}
	return
}
