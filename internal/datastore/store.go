package datastore

import "github.com/SimplQ/simplQ-golang/internal/models/db"

type QueueStore interface {
	// Create a new queue and return the queue ID.
	CreateQueue(db.Queue) db.QueueId

	// Read a queue by id.
	ReadQueue(db.QueueId) db.Queue

	// Set the queue pause status to true
	PauseQueue(db.QueueId)

    // Set the queue pause status to false
	ResumeQueue(db.QueueId)
	
    // Set the queue delete status to new value.
	DeleteQueue(db.QueueId)

	// Add a new token to the queue.
	AddTokenToQueue(db.QueueId, db.Token)

	// Read token by id.
	ReadToken(db.TokenId)

	// Delete token
	RemoveToken(db.TokenId)
}

var Store QueueStore = NewMongoDB()
