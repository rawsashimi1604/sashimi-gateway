package main

import (
	"log"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/tuna-api/api"
)

func main() {
	r := api.NewRouter()
	log.Println("Tuna Server is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
