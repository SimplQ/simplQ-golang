// Package internal/models/api defines models to be used for API requests.
package api

import (
	"fmt"
	"log"

	"github.com/SimplQ/simplQ-golang/internal/models/db"
	"github.com/ttacon/libphonenumber"
)

// AddTokenRequest is a model to structure an add token request.
//
// Members
//
// QueueId          -   QueueId of the queue that this token is to be added to
// Name             -   Name of the token, typically name of the person whom the
//                      token was issued to.
// ContactNumber    -   Contact Number.
// EmailId          -   Optional. Email ID if the queue collects email ID of users.
type AddTokenRequest struct {
	QueueId       db.QueueId
	Name          string
	ContactNumber string
	EmailId       string
}

// AddTokenResponse is a model to structure the response of an add token request.
type AddTokenResponse db.TokenId

// Minimum length of a token name
const MIN_LENGTH_TOKEN_NAME = 4

// Maximum length of a token name
const MAX_LENGTH_TOKEN_NAME = 20

func (req AddTokenRequest) Validate() (ValidationError, bool) {
	if len(req.Name) < MIN_LENGTH_TOKEN_NAME || len(req.Name) > MAX_LENGTH_TOKEN_NAME {
		message := fmt.Sprintf("Token name length should be greater than %d characters and less than %d characters", MIN_LENGTH_TOKEN_NAME, MAX_LENGTH_TOKEN_NAME)
		return ValidationError{
			Field:   "Name",
			Message: message,
		}, false
	} else if num, err := libphonenumber.Parse(req.ContactNumber, "IN"); err != nil {
		log.Println(num)
		message := fmt.Sprint(err)
		return ValidationError{
			Field:   "ContactNumber",
			Message: message,
		}, false
	} else {
		return ValidationError{}, true
	}
}
