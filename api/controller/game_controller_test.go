package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to handler
	req, err := http.NewRequest("GET", "/api/health-check", nil)
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record response
	recordResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	// Call the ServeHTTP method to directly pass in our request and ResponseRecorder
	handler.ServeHTTP(recordResponse, req)

	// Check to see if it sends back correct status
	if status := recordResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned back wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	// Check to see if response body is what we expected
	expected := `{"alive":true}`
	if responseBody := recordResponse.Body.String(); responseBody != expected {
		t.Errorf("handler returned back unexpected body: got %v want %v\n", responseBody, expected)
	}
}
