package controller

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler function
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"alive":true}`)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
