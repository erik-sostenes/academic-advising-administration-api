package services

import (
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Encrytor structure  that contains the methods for encrypting password
type Encrytor struct {}

// EncryptPassword will encrypt the password
func (e *Encrytor) EncryptPassword(password string)(string, error){
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
	return string(encryptPassword), model.InternalServerError(err.Error())
}

// ValidatePassword validate the encrypted password if correct
func (e *Encrytor) ValidatePassword(hashedPassword, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err = model.Forbidden("Your password is forbidden.")
		return
	}
	return
}
