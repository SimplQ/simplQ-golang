package db

import (
	"time"
)

// This ID will be exposed to clients, and hence have to be properly random and unguessable.
type TokenId Id

type Token struct {
	// Unique ID for the token.
	Id TokenId `bson:"_id,omitempty"`

	// QueueId of the Queue that this token belongs to
	QueueId QueueId `bson:"queueId"`

	// Number representing the position of the token in the queue
	TokenNumber uint32 `bson:"tokenNumber"`

	// Name of the token, typically name of the person to whom the token was
	// issued.
	Name string `bson:"name"`

	// Contact Number
	ContactNumber string `bson:"contactNumber"`

	// Optional. Email ID if the queue collects email ID of users.
	EmailId string `bson:"emailId"`

	// Set to true if the token has been deleted
	IsDeleted bool `bson:"isDeleted"`

	// Number of times the token was notified.
	NotifiedCount uint32 `bson:"notifiedCount"`

	// Timestamp when the token was last notified.
	LastNotifiedTime time.Time `bson:"lastNotifiedTime,omitempty"`

	// Creation time.
	CreationTime time.Time `bson:"creationTime,omitempty"`

	// Deletion time.
	DeletionTime time.Time `bson:"deletionTime,omitempty"`
}
