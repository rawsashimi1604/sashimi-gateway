package main

import (
	"log"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/salmon-api/api"
)

func main() {
	r := api.NewRouter()
	log.Println("Salmon Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
