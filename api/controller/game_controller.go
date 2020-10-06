package controller

import (
	"net/http"
)

// Test handler function
func testConnection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Connected"))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
