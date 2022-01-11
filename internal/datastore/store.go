package datastore

import (
	"github.com/SimplQ/simplQ-golang/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(models.Queue) models.QueueId

	// Read a queue by id.
	ReadQueue(models.QueueId) (models.Queue, error)

	// Set the queue pause status to true/false
	SetIsPaused(models.QueueId, bool) (*mongo.UpdateResult, error)

	// Set the queue delete status to new value.
	DeleteQueue(models.QueueId) (*mongo.DeleteResult, error)

	// Add a new token to the queue.
	AddTokenToQueue(models.QueueId, models.Token)

	// Read token by id.
	ReadToken(models.TokenId)

	// Delete token
	RemoveToken(models.TokenId)
}

var Store QueueStore = NewMongoDB()
