package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models"
)

func Queue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getQueue(w, r)
	case "POST":
		createQueue(w, r)
	case "DELETE":
		deleteQueue(w, r)
	default:
		fmt.Fprintf(w, "Invalid method")
	}
}

func QueuePause(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pauseQueue(w, r)
	} else {
		fmt.Fprintf(w, "Invalid method")
	}
}

func QueueResume(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		resumeQueue(w, r)
	} else {
		fmt.Fprintf(w, "Invalid method")
	}
}

func getQueue(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "GET Queue not implemented")
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Get Id from path
	id := parts[3]
	log.Println("URL path param 'id' is: " + string(id))
	q := datastore.Store.ReadQueue(models.QueueId(id))
	json.NewEncoder(w).Encode(q)
	fmt.Fprintf(w, "Get queue")
}

func createQueue(w http.ResponseWriter, r *http.Request) {
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

func pauseQueue(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Get Id from path
	id := parts[4]
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.PauseQueue(models.QueueId(id))
	fmt.Fprintf(w, "Queue paused")
}

func resumeQueue(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//Get Id from path
	id := parts[4]
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.ResumeQueue(models.QueueId(id))
	fmt.Fprintf(w, "Queue resumed")
}

func deleteQueue(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Get Id from path
	id := parts[3]
	log.Println("URL path param 'id' is: " + string(id))
	datastore.Store.DeleteQueue(models.QueueId(id))
	fmt.Fprintf(w, "Delete queue")
}
