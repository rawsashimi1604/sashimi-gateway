package api

import (
	"github.com/gorilla/mux"
)

// New creates a new router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", GetDishesHandler).Methods("GET")
	r.HandleFunc("/", AddDishHandler).Methods("POST")
	r.HandleFunc("/test", TestDish).Methods("GET")
	return r
}
