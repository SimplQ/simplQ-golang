package datastore

import (
	"github.com/SimplQ/simplQ-golang/internal/models/db"
)

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(db.Queue) (db.QueueId, error)

	// Read a queue by id.
	ReadQueue(db.QueueId) (db.Queue, error)

	// Set the queue pause status to true/false
	SetIsPaused(db.QueueId, bool) error

	// Set the queue delete status to new value.
	DeleteQueue(db.QueueId) error

	// Add a new token to the queue.
	AddTokenToQueue(db.QueueId, db.Token)

	// Read token by id.
	ReadToken(db.TokenId)

	// Delete token
	RemoveToken(db.TokenId)
}

var Store QueueStore = NewMongoDB()
