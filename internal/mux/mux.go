package mux

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/SimplQ/simplQ-golang/internal/handler"
)

func InitalizeRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(10 * time.Second))

	// Routes for "queue" resource
	r.Route("/queue", func(r chi.Router) {
		// POST /articles
		r.Post("/", handler.CreateQueue)

		// Subrouters
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.QueueCtx)
			r.Get("/", handler.GetQueue)       // GET /queue/123
			r.Delete("/", handler.DeleteQueue) // DELETE /queue/123
		})
		r.Route("/pause/{id}", func(r chi.Router) {
			r.Use(handler.QueueCtx)
			r.Put("/", handler.PauseQueue) // PUT /queue/pause/123

		})
		r.Route("/resume/{id}", func(r chi.Router) {
			r.Use(handler.QueueCtx)
			r.Put("/", handler.ResumeQueue) // PUT /queue/resume/123
		})
	})
	return r
}
