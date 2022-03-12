package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models/api"
	"github.com/SimplQ/simplQ-golang/internal/models/db"
	"github.com/go-chi/chi/v5"
)

type key int

const queueId key = 0

func GetQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	uid := r.Context().Value("uid").(string)
	log.Println(uid)

	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusBadRequest)
		return
	}

	queue, err := datastore.Store.ReadQueue(db.QueueId(id))

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(queue)
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var queueRequest api.CreateQueueRequest
	err := decoder.Decode(&queueRequest)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
	}

	// Validation
	validationErr, ok := queueRequest.Validate()

	if !ok {
		http.Error(w, validationErr.Message, http.StatusBadRequest)
		return
	}

	// Initialize values
	// Only consider queue name from the body of the request
	queue := db.Queue{
		QueueName:    queueRequest.QueueName,
		CreationTime: time.Now(),
		IsDeleted:    false,
		IsPaused:     false,
		Tokens:       make([]db.Token, 0),
	}

	insertedId, err := datastore.Store.CreateQueue(queue)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	log.Printf("Inserted %s", insertedId)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Queue created with Id: %s", insertedId)
}

func PauseQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusBadRequest)
		return
	}
	err := datastore.Store.SetIsPaused(db.QueueId(id), true) // Set IsPaused = true
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Paused Queue Id: %s", id)
	w.WriteHeader(http.StatusOK)
}

func ResumeQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusBadRequest)
		return
	}
	err := datastore.Store.SetIsPaused(db.QueueId(id), false) // Set IsPaused = false
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Resumed Queue Id: %s", id)
	w.WriteHeader(http.StatusOK)
}

func DeleteQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(queueId).(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusBadRequest)
		return
	}
	err := datastore.Store.DeleteQueue(db.QueueId(id))
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
