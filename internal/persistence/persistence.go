package persistence

import (
	"github.com/SimplQ/simplQ-golang/internal/models"
)

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(models.Queue) models.QueueId

	// Read a queue by id.
	ReadQueue(models.QueueId) models.Queue

	// Set the queue pause status to new value.
	UpdateQueuePauseStatus(models.QueueId, bool)

	// Set the queue delete status to new value.
	UpdateQueueDeleteStatus(models.QueueId, bool)

	// Add a new token to the queue.
	AddTokenToQueue(models.QueueId, models.Token)

	// Read token by id.
	ReadToken(models.TokenId)

	// Set token status to new value.
	UpdateTokenDeleteStatus(models.TokenId, bool)
}
