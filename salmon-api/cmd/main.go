package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const SalmonType string = "Salmon"

// Define a simple struct for a Salmon Dish entity
type Salmon struct {
	Id          string `json:"id"`
	Item        string `json:"item"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Create some sample data
var dishes = []Salmon{
	{Id: "1", Item: "Salmon Nigiri Sushi", Description: "Salmon Nigiri Sushi comes from Japan and is a staple sushi food.", Type: SalmonType},
	{Id: "2", Item: "Salmon Don", Description: "Delicious and tasty dish: salmon with rice.", Type: SalmonType},
}

// GetDishesHandler returns all salmon dishes
func GetDishesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Attach the handlers
	r.HandleFunc("/", GetDishesHandler).Methods("GET")

	// Start the server
	log.Println("Salmon Server is running on port 8001...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
