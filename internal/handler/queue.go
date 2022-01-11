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

func GetQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), 404)
		return
	}
	queue, err := datastore.Store.ReadQueue(models.QueueId(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), 404)
		return
	}
	json.NewEncoder(w).Encode(queue)
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var q models.Queue
	err := decoder.Decode(&q)

	if err != nil {
		panic(err)
	}

	// Initialize values
	// Only consider queue name from the body of the request
	q.CreationTime = time.Now()
	q.IsDeleted = false
	q.IsPaused = false
	q.Tokens = make([]models.Token, 0)

	log.Print("Create Queue: ")
	log.Println(q)

	insertedId := datastore.Store.CreateQueue(q)

	log.Printf("Inserted %s", insertedId)

	fmt.Fprintf(w, "Post queue")
}

func PauseQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), 404)
		return
	}
	result, err := datastore.Store.SetIsPaused(models.QueueId(id), true) // Set IsPaused = true
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), 404)
		return
	}
	fmt.Fprintf(w, "Paused Queue Id: %s", id)
	w.WriteHeader(200)
}

func ResumeQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), 404)
		return
	}
	result, err := datastore.Store.SetIsPaused(models.QueueId(id), false) // Set IsPaused = false
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, fmt.Sprintf("No record found for Queue Id: %s", id), 404)
		return
	}
	fmt.Fprintf(w, "Resumed Queue Id: %s", id)
	w.WriteHeader(200)
}

func DeleteQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.DeleteQueue(models.QueueId(id))
	fmt.Fprintf(w, "Delete queue")
}

func QueueCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
