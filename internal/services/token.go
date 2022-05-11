package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/itsoeh/academic-advising-administration-api/internal/model"
)

// Token structure whose task is to generate and validate a token
type Token struct{}

// GeterateToken method that will generate the JSON web token
func (*Token) GeterateToken(email string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims {
			"iss": "itsoeh.edu.mx",
			"aud": email,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err = token.SignedString(singKey)
	if err != nil {
		err = model.InternalServerError("The token could not be generate.")
		return
	}
	return
}

// ValidateToken validate and return a token
func(t *Token) ValidateToken(tokenString string) (err error) {
	token, err := t.parseToken(tokenString)
	if err != nil {
		return 
	}

	if !token.Valid {
		err = errors.New("Does not have authorization.")
		return 
	}
	return 
}

// parseToken will receive the parsed token and should return the key for validating
func (*Token) parseToken(tokenString string) (*jwt.Token, error) {
		return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, errors.New("Invalid signing method.")
			}

			iss := "itsoeh.edu.mx"
			verifyIss := t.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !verifyIss {
				return nil, errors.New("Invalid issuer.")
			}
			
			return verifyKey, nil
		})
}
