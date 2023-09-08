package api

import (
	"github.com/gorilla/mux"
)

// New creates a new router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", Healthz).Methods("GET")
	r.HandleFunc("/", GetDishesHandler).Methods("GET")
	return r
}
