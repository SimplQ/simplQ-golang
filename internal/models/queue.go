package models

import (
	"time"
)

// This ID will be exposed to clients, and hence have to be properly random and unguessable.
type QueueId Id

type Queue struct {
	// Unique ID for the queue. 
	Id QueueId

	// Name of the queue. The name is used in generating queue urls, for ex 
	// https://simplq.me/j/<QueueName>. The storage layer guarantees sure that 
	// there is only one queue by a given name.
	QueueName string

	// Set to true if the queue is temporarily not issuing tokens
	bool IsPaused

	// Set to true if the queue has been deleted
	bool IsDeleted

	// Tokens present in the queue.
	Tokens []Token

	// Creation time.
	CreationTime time.Time

	// Deletion time.
	DeletionTime time.Time
}