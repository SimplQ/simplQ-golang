package handler

import (
	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/persistence"
)

var Store persistence.QueueStore = datastore.NewMongoDB()
