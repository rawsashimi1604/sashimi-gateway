package api

import (
	"github.com/gorilla/mux"
)

// New creates a new router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", GetDishesHandler).Methods("GET")
	r.HandleFunc("/", AddDishHandler).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", GetDishByIdHandler).Methods("GET")
	r.HandleFunc("/test", TestDish).Methods("GET")
	r.HandleFunc("/test-salmon", TestDish).Methods("GET")
	return r
}
