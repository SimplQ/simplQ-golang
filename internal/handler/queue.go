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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	q := datastore.Store.ReadQueue(models.QueueId(id))
	json.NewEncoder(w).Encode(q)
	fmt.Fprintf(w, "Get queue")
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.PauseQueue(models.QueueId(id))
	fmt.Fprintf(w, "Queue paused")
}

func ResumeQueue(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.ResumeQueue(models.QueueId(id))
	fmt.Fprintf(w, "Queue resumed")
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
