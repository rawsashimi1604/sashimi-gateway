package model

const TunaType string = "Tuna"

// Tuna represents the model for a Tuna dish entity
type Tuna struct {
	Id          string `json:"id"`
	Item        string `json:"item"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
