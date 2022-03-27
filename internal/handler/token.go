package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimplQ/simplQ-golang/internal/authentication"
	"github.com/SimplQ/simplQ-golang/internal/datastore"
	"github.com/SimplQ/simplQ-golang/internal/models/api"
	"github.com/SimplQ/simplQ-golang/internal/models/db"
	"github.com/go-chi/chi/v5"
)

const TOKEN_ID = "tokenId"

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
	id := r.Context().Value(TOKEN_ID).(string)

	if id == "" {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	token, err := datastore.Store.ReadToken(db.TokenId(id))

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete token")
	id := r.Context().Value(TOKEN_ID).(string)
	uid := r.Context().Value(authentication.UID).(string)

	if id == "" {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	token, err := datastore.Store.ReadToken(db.TokenId(id))

	log.Println("Read token")

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	queueId := token.QueueId

	queue, err := datastore.Store.ReadQueue(db.QueueId(queueId))

	if queue.Owner != uid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = datastore.Store.RemoveToken(db.TokenId(id))

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted token %s", id)
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), TOKEN_ID, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
