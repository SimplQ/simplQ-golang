package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models/db"
	"github.com/SimplQ/simplQ-golang/internal/models/api"
)

func GetQueue(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "GET Queue not implemented")
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var q api.CreateQueueRequest
    err := decoder.Decode(&q)

    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
    }

    // Validation
    validation_err, ok := q.Validate() 

    if !ok {
        http.Error(w, validation_err.Message, http.StatusBadRequest)
    }

    // Initialize values
    // Only consider queue name from the body of the request
    queue := db.Queue {
        QueueName: q.QueueName,
        CreationTime: time.Now(),
        IsDeleted: false,
        IsPaused: false,
        Tokens: make([]db.Token, 0),
    }

    log.Print("Create Queue: ")
    log.Println(q)

    insertedId := datastore.Store.CreateQueue(queue)

    log.Printf("Inserted %s", insertedId)

    fmt.Fprintf(w, "Post queue")
}
