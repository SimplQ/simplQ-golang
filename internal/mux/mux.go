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
			r.Get("/", handler.GetQueue)          // GET /queue/123
			r.Put("/pause", handler.PauseQueue)   // PUT /queue/123/pause
			r.Put("/resume", handler.ResumeQueue) // PUT /queue/123/resume
			r.Delete("/", handler.DeleteQueue)    // DELETE /queue/123
		})
	})

	// Routes for "token" resource
	r.Route("/token", func(r chi.Router) {
		// Add new token to queue
		r.Post("/", handler.CreateToken)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.TokenCtx)
			r.Get("/", handler.GetToken)
		})
	})
	return r
}
