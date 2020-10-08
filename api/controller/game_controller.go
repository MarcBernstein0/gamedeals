package controller

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler function
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"alive":true}`)
}
