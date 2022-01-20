// Package internal/models/api defines models to be used for API requests.
package api

// Structure ValidationError is used to describe an error that occurs in a
// validation.
//
// Members
//
// Field    -   name of the field which had an error.
// Message  -   description of the error.
type ValidationError struct {
	Field   string
	Message string
}

// Interface Validator defines how to define a validator for a request.
type Validator interface {
    // Method Validate validates the request data.
    // Returns ValidationError, false if the request data is invalid.
    // Returns ValidationError, true if the request data is valid.
    // If ValidationError, true is returned ValidationError is empty and 
    // should be ignored.
	Validate() (ValidationError, bool)
}
