package mux

import (
	"net/http"

	"github.com/SimplQ/simplQ-golang/internal/handler"
)

func InitalizeRoutes() {
	// Both paths are needed since /api/queue/ doesn't cover /api/queue
	http.HandleFunc("/api/queue/", handler.Queue)
	http.HandleFunc("/api/queue", handler.Queue)
	http.HandleFunc("/api/queue/pause/", handler.QueuePause)
	http.HandleFunc("/api/queue/resume/", handler.QueueResume)
}
