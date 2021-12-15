package persistence

import (
	"github.com/SimplQ/simplQ-golang/internal/models"
)

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(models.Queue) models.QueueId

	// Read a queue by id.
	ReadQueue(models.QueueId) models.Queue

	// Set the queue pause status to true
	PauseQueue(models.QueueId)

    // Set the queue pause status to false
	ResumeQueue(models.QueueId)
	
    // Set the queue delete status to new value.
	DeleteQueue(models.QueueId)

	// Add a new token to the queue.
	AddTokenToQueue(models.QueueId, models.Token)

	// Read token by id.
	ReadToken(models.TokenId)

	// Set token status to new value.
	UpdateTokenDeleteStatus(models.TokenId, bool)
}
