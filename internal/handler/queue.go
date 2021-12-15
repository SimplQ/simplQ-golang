package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
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
    w.Header().Add("Content-Type", "application/json")
    encoder := json.NewEncoder(w)
    q := models.Queue{
        QueueName: "alice",
    }
    encoder.Encode(q)
}

func createQueue(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var q models.Queue
    err := decoder.Decode(&q)

    if err != nil {
        panic(err)
    }

    fmt.Print("Create Queue: ")
    fmt.Println(q)

    fmt.Fprintf(w, "Post queue")
}
