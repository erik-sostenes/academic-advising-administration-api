package services

import "golang.org/x/crypto/bcrypt"

// Encrytor structure  that contains the methods for encrypting password
type Encrytor struct {}

// EncryptPassword will encrypt the password
func (e *Encrytor) EncryptPassword(password string)(string, error){
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
	return string(encryptPassword), err
}
