package api

import (
	"encoding/json"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/tuna-api/model"
)

var dishes = []model.Tuna{
	{Id: "1", Item: "Otoro", Description: "Fatty tuna, super delicious.", Type: model.TunaType},
	{Id: "2", Item: "Chirashi", Description: "Tuna sashimi served with rice and wasabi.", Type: model.TunaType},
}

// GetDishesHandler retrieves all salmon dishes
func GetDishesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}
