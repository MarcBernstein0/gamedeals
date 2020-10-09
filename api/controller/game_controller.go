package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/middleware"
)

// healthCheckHandler function
func healthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, `{"alive":true}`)
}

func findGameHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// TODO: Make post request
		rw.WriteHeader(http.StatusAccepted)
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SetupRoutes function
func SetupRoutes(apiPath string) {
	healthHandler := http.HandlerFunc(healthCheckHandler)
	http.Handle(fmt.Sprintf("%s/health-check", apiPath), middleware.HeaderMiddleWare(healthHandler))
}
