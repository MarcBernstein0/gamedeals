package controller

import (
	"bytes"
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

func TestGameHandler(t *testing.T) {

	jsonRequest := []byte(`{"title":"string","system":"PC"}`)

	log.Println("Test 1: Test Wrong Method")
	req1, err := http.NewRequest("GET", "/api/findGame", nil)
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse1 := sendRequest(req1, findGameHandler)

	if status := recordResponse1.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Method sent was not allowed but did not return a MethodNotAllowed(405) error\nexpected: %d got: %d\n", http.StatusMethodNotAllowed, status)
	}

	log.Println("Test 2: Test Status Response")

	req2, err := http.NewRequest("POST", "/api/findGame", bytes.NewBuffer(jsonRequest))
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse2 := sendRequest(req2, findGameHandler)

	if status := recordResponse2.Code; status != http.StatusAccepted {
		t.Errorf("handler returned back wrong status code: got %v want %v\n", status, http.StatusAccepted)
	}

	log.Println("Test 3: Test nil Post")
	req3, err := http.NewRequest("POST", "/api/findGame", nil)
	// log.Println(req3)
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse3 := sendRequest(req3, findGameHandler)
	if status := recordResponse3.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned back wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}

	log.Println("Test 4: Handle Non-Nil Posts Bad Request")
	badRequest := []byte(`{"title": "string"`)
	req4, err := http.NewRequest("POST", "/api/findGame", bytes.NewBuffer(badRequest))
	if err != nil {
		log.Fatalf("Error when creating test request\n%v\n", err)
	}

	recordResponse4 := sendRequest(req4, findGameHandler)
	log.Println(recordResponse4)
	if status := recordResponse4.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned back wrong status code: got %v want %v\n", status, http.StatusBadRequest)
	}

}

func sendRequest(req *http.Request, handleFunc func(rw http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	recorderResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(handleFunc)
	handler.ServeHTTP(recorderResponse, req)
	return recorderResponse
}
