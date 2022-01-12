package api

import (
    "fmt"

    "github.com/SimplQ/simplQ-golang/internal/models/db"
)

type CreateQueueRequest struct {
    QueueName string
}

type CreateQueueResponse db.Queue

const MIN_LENGTH = 4
const MAX_LENGTH = 20

func (req CreateQueueRequest) Validate() (ValidationError, bool) {
    if len(req.QueueName) < MIN_LENGTH || len(req.QueueName) > MAX_LENGTH {
        message := fmt.Sprintf("Queue name length should be greater than %d characters and less than %d charaacters", MIN_LENGTH, MAX_LENGTH)
        return ValidationError {
            Field: "QueueName",
            Message: message,
        }, false
    } 

    return ValidationError{}, true
}
