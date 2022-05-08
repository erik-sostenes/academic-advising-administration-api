package services

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/itsoeh/academy-advising-administration-api/internal/model"
)

// Certifier will have the method for upload certificates and parse with JWT
type Certifier interface{
	// Certificates is a singleton method to load the certificates
	Certificates(publicCertificate, privateCertificate string) error
}

// variables
var (
	syncOnce sync.Once
	singKey *rsa.PrivateKey
	verifyKey * rsa.PublicKey
)

type certifier struct {}

func NewCertifier() Certifier {
	return &certifier{}
}

func (c *certifier) Certificates(publicCertificate, privateCertificate string) (err error) {
	syncOnce.Do(func() {
		err = c.certificates(publicCertificate, privateCertificate)
	})
	return 
}

// certificates load the certificates
func (c *certifier) certificates(publicCertificate, privateCertificate string) (err error) {
publicBytes, err := ioutil.ReadFile(publicCertificate)
	if err != nil {
		err = model.InternalServerError("The public certificate not fount.")
		return
	}
	privateBytes, err := ioutil.ReadFile(privateCertificate)
	if err != nil {
		err = model.InternalServerError("The private certificate not fount.")
		return
	}

	return c.parseRSA(publicBytes, privateBytes)
}

// parseRSA parse the certificates with JWT
func (c *certifier) parseRSA(public, private []byte) (err error) {
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return
	}

	singKey, err = jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return
	}
	return
}
