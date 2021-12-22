package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/models"
)

func Queue(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            getQueue(w, r)
        case "POST":
            createQueue(w, r)
        default:
            fmt.Fprintf(w, "Invalid method")
    }
}

func getQueue(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "GET Queue not implemented")
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

    insertedId := Store.CreateQueue(q)

    log.Printf("Inserted %s", insertedId)

    fmt.Fprintf(w, "Post queue")
}
