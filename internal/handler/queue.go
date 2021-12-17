package handler

import (
	"encoding/json"
	"log"
    "fmt"
	"net/http"

	"github.com/SimplQ/simplQ-golang/internal/models"
	"github.com/SimplQ/simplQ-golang/internal/persistence"
)

var Store persistence.QueueStore

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

    log.Print("Create Queue: ")
    log.Println(q)

    insertedId := Store.CreateQueue(q)

    log.Printf("Inserted %s", insertedId)

    fmt.Fprintf(w, "Post queue")
}
