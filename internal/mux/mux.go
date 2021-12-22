package mux

import (
	"net/http"

	"github.com/SimplQ/simplQ-golang/internal/handler"
	"github.com/SimplQ/simplQ-golang/internal/persistence"
)

func InitalizeRoutes(store persistence.QueueStore) {
    // Both paths are needed since /api/queue/ doesn't cover /api/queue
    http.HandleFunc("/api/queue/", handler.Queue)
    http.HandleFunc("/api/queue", handler.Queue)
}
