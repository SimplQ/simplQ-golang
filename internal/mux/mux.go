package mux

import (
	"net/http"

	"github.com/SimplQ/simplQ-golang/internal/handler"
	"github.com/SimplQ/simplQ-golang/internal/persistence"
)

func InitalizeRoutes(store persistence.QueueStore) {
    // Set the QueueStore in the handler package
    handler.Store = store

    // Both paths are needed since /api/queue/ doesn't cover /api/queue
    http.HandleFunc("/api/queue/", handler.Queue)
    http.HandleFunc("/api/queue", handler.Queue)
}
