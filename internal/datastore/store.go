package datastore

import "github.com/SimplQ/simplQ-golang/internal/persistence"

var Store persistence.QueueStore = NewMongoDB()
