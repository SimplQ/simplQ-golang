package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SimplQ/simplQ-golang/internal/mux"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables if .env file is present
    // ignore if not present
    godotenv.Load();

	routes := mux.InitalizeRoutes()

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, routes))
}
