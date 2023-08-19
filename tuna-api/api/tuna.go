package api

import (
	"encoding/json"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/salmon-api/model"
)

var dishes = []model.Salmon{
	{Id: "1", Item: "Salmon Nigiri Sushi", Description: "Salmon Nigiri Sushi comes from Japan and is a staple sushi food.", Type: model.SalmonType},
	{Id: "2", Item: "Salmon Don", Description: "Delicious and tasty dish: salmon with rice.", Type: model.SalmonType},
}

// GetDishesHandler retrieves all salmon dishes
func GetDishesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}
