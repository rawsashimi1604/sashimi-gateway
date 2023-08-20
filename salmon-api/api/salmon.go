package api

import (
	"encoding/json"
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/salmon-api/model"
	"github.com/rawsashimi1604/sashimi-gateway/salmon-api/utils"
)

var dishes = []model.Salmon{
	{Id: 1, Item: "Salmon Nigiri Sushi", Description: "Salmon Nigiri Sushi comes from Japan and is a staple sushi food.", Type: model.SalmonType},
	{Id: 2, Item: "Salmon Don", Description: "Delicious and tasty dish: salmon with rice.", Type: model.SalmonType},
}

// GetDishesHandler retrieves all salmon dishes
func GetDishesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}

// AddDishHandler adds a new salmon object to the list of dishes
func AddDishHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSalmon = model.Salmon{}
	err := json.NewDecoder(r.Body).Decode(&newSalmon)
	if err != nil {
		utils.JSONError(w, "failed to decode json body", http.StatusBadRequest)
		return
	}

	if newSalmon.Id == 0 || newSalmon.Item == "" || newSalmon.Description == "" {
		utils.JSONError(w, "required fields are missing", http.StatusBadRequest)
		return
	}
	newSalmon.Type = model.SalmonType
	dishes = append(dishes, newSalmon)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSalmon)
}
