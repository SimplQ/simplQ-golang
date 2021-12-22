package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/mux"
)

func main() {
	listenAddr := ":8080"

    // Use local mongodb instance if env variable not set
    mongoUri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
    
    if val, ok := os.LookupEnv("MONGO_URI"); ok {
        mongoUri = val
	}

    mongodb := datastore.NewMongoDB(mongoUri)

    mux.InitalizeRoutes(mongodb)

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
