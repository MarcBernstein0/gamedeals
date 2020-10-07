package controller

import (
	"net/http"
)

// TestHandler function
func TestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Connected"))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
