// Package internal/models/api defines models to be used for API requests.
package api

import (
	"fmt"

	"github.com/SimplQ/simplQ-golang/internal/models/db"
)

// CreateQueueRequest is a model to structure a create queue request.
//
// Members
//
// QueueName    -   Name of the queue to be created.
type CreateQueueRequest struct {
	QueueName string
}

// CreateQueueResponse is a model to strcuture the response of a create queue
// request.
type CreateQueueResponse db.Queue

// Minimum length of a queue name.
const MIN_LENGTH = 4

// Maximum length of a queue name.
const MAX_LENGTH = 20

// Validate function for CreateQueueRequest validates if the queue name is within
// the defined range.
func (req CreateQueueRequest) Validate() (ValidationError, bool) {
	if len(req.QueueName) < MIN_LENGTH || len(req.QueueName) > MAX_LENGTH {
		message := fmt.Sprintf("Queue name length should be greater than %d characters and less than %d charaacters", MIN_LENGTH, MAX_LENGTH)
		return ValidationError{
			Field:   "QueueName",
			Message: message,
		}, false
	}

	return ValidationError{}, true
}
