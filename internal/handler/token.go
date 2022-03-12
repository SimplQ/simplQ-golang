package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models/api"
	"github.com/SimplQ/simplQ-golang/internal/models/db"
	"github.com/go-chi/chi/v5"
)

const tokenId key = 0

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

func GetToken(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(tokenId).(string)

	if id == "" {
		http.Error(w, fmt.Sprintf("Invalid Id: %s", id), http.StatusBadRequest)
		return
	}

	token, err := datastore.Store.ReadToken(db.TokenId(id))

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}

func TokenCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), tokenId, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
