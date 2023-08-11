package api

import (
	"fmt"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	fmt.Println("Creating Mux Router.")

	router := mux.NewRouter()
	return router
}
