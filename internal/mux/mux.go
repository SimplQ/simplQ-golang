package mux

import (
    "net/http"
    "github.com/SimplQ/simplQ-golang/internal/handler"
)

func InitalizeRoutes() {
	http.HandleFunc("/api/HttpExample", handler.HelloHandler);
}
