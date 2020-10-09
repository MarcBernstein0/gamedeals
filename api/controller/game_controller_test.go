package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
	handler := http.HandlerFunc(healthCheckHandler)

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

func TestGameHandlerBadMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/findGame", nil)
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(findGameHandler)

	handler.ServeHTTP(recordResponse, req)

	if status := recordResponse.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Method sent was not allowed but did not return a MethodNotAllowed(405) error\nexpected: %d got: %d\n", http.StatusMethodNotAllowed, status)
	}
}

func TestGameHandler(t *testing.T) {
	jsonRequest := strings.NewReader(`{"title":"string","system":"PC"}`)
	req, err := http.NewRequest("POST", "/api/findGame", jsonRequest)
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(findGameHandler)

	handler.ServeHTTP(recordResponse, req)

	if status := recordResponse.Code; status != http.StatusAccepted {
		t.Errorf("handler returned back wrong status code: got %v want %v\n", status, http.StatusAccepted)
	}
}
