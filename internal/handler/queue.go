package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models"
	"github.com/go-chi/chi/v5"
)

type key int

const queueId key = 0

func GetQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusNotFound)
		return
	}
	queue, err := datastore.Store.ReadQueue(models.QueueId(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(queue)
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var q models.Queue
	err := decoder.Decode(&q)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
	}

	// Initialize values
	// Only consider queue name from the body of the request
	q.CreationTime = time.Now()
	q.IsDeleted = false
	q.IsPaused = false
	q.Tokens = make([]models.Token, 0)

	insertedId, err := datastore.Store.CreateQueue(q)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	log.Printf("Inserted %s", insertedId)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Queue created with Id: %s", insertedId)
}

func PauseQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusNotFound)
		return
	}
	result, err := datastore.Store.SetIsPaused(models.QueueId(id), true) // Set IsPaused = true
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Paused Queue Id: %s", id)
	w.WriteHeader(http.StatusOK)
}

func ResumeQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusNotFound)
		return
	}
	result, err := datastore.Store.SetIsPaused(models.QueueId(id), false) // Set IsPaused = false
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Resumed Queue Id: %s", id)
	w.WriteHeader(http.StatusOK)
}

func DeleteQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusNotFound)
		return
	}
	result, err := datastore.Store.DeleteQueue(models.QueueId(id))
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Deleted Queue Id: %s", id)
	w.WriteHeader(http.StatusOK)
}

func QueueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), queueId, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
