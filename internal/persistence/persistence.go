package persistence

import (
	"github.com/SimplQ/simplQ-golang/internal/models"
)

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(models.Queue) models.QueueId
	
	// Read a queue by id.
	ReadQueue(models.QueueId) models.Queue

	// Set the queue status to new value.
	UpdateQueueStatus(models.QueueId, models.QueueStatus)

	// Add a new token to the queue.
	AddTokenToQueue(models.QueueId, models.Token)

	// Read token by id.
	ReadToken(models.TokenId)

	// Set token status to new value.
	UpdateTokenStatus(models.TokenId, models.TokenStatus)
}