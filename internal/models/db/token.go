package db

import (
	"time"
)

// This ID will be exposed to clients, and hence have to be properly random and unguessable.
type TokenId Id

type Token struct {
	// Unique ID for the token.
	Id TokenId

	// QueueId of the Queue that this token belongs to
	QueueId QueueId

	// Name of the token, typically name of the person to whom the token was
	// issued.
	Name string

	// Contact Number
	ContactNumber string

	// Optional. Email ID if the queue collects email ID of users.
	EmailId string

	// Set to true if the token has been deleted
	IsDeleted bool

	// Number of times the token was notified.
	NotifiedCount uint32

	// Timestamp when the token was last notified.
	LastNotifiedTime time.Time

	// Creation time.
	CreationTime time.Time

	// Deletion time.
	DeletionTime time.Time
}
