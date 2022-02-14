package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models/api"
	"github.com/SimplQ/simplQ-golang/internal/models/db"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var tokenRequest api.AddTokenRequest
	err := decoder.Decode(&tokenRequest)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	validationErr, ok := tokenRequest.Validate()

	if !ok {
		http.Error(w, validationErr.Message, http.StatusBadRequest)
		return
	}

	token := db.Token{
		Name:          tokenRequest.Name,
		ContactNumber: tokenRequest.ContactNumber,
		EmailId:       tokenRequest.EmailId,
		IsDeleted:     false,
		NotifiedCount: 0,
		CreationTime:  time.Now(),
	}

	insertedId, err := datastore.Store.AddTokenToQueue(tokenRequest.QueueId, token)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	log.Printf("Inserted %s", insertedId)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Token created with Id: %s", insertedId)
}
