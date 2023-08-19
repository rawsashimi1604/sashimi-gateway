package model

const SalmonType string = "Salmon"

// Salmon represents the model for a salmon dish entity
type Salmon struct {
	Id          string `json:"id"`
	Item        string `json:"item"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
