package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONError(t *testing.T) {
	// Define test error message and status code
	errorMessage := "This is a test error"
	statusCode := http.StatusBadRequest

	// Create a mock request and response writer
	_, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Call the JSONError function
	JSONError(rr, errorMessage, statusCode)

	// Check status code
	if status := rr.Code; status != statusCode {
		t.Errorf("JSONError returned wrong status code: got %v want %v", status, statusCode)
	}

	// Decode the returned body to check if it contains the error message
	var returnedError JSONErrorMessage
	err = json.NewDecoder(rr.Body).Decode(&returnedError)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}

	if returnedError.Message != errorMessage {
		t.Errorf("JSONError returned unexpected body: got %s want %s", returnedError.Message, errorMessage)
	}
}
