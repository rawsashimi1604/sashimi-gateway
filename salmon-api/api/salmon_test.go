package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rawsashimi1604/sashimi-gateway/salmon-api/model"
)

func TestGetDishesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDishesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAddDishHandler(t *testing.T) {
	newSalmon := model.Salmon{
		Id:          3,
		Item:        "Test Salmon",
		Description: "This is a test salmon dish.",
	}
	newSalmonBytes, err := json.Marshal(newSalmon)
	if err != nil {
		t.Fatal("Failed to marshal test salmon dish")
	}

	req, err := http.NewRequest("POST", "/add-dish", bytes.NewReader(newSalmonBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddDishHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decode the returned body to check if it contains the newSalmon
	var returnedSalmon model.Salmon
	err = json.NewDecoder(rr.Body).Decode(&returnedSalmon)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if returnedSalmon.Id != newSalmon.Id || returnedSalmon.Item != newSalmon.Item || returnedSalmon.Description != newSalmon.Description {
		t.Errorf("handler returned unexpected body: got %+v want %+v", returnedSalmon, newSalmon)
	}
}
