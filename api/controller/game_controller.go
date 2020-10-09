package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/model"
)

// healthCheckHandler function
func healthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, `{"alive":true}`)
}

func findGameHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		// check if body is empty
		if r.Body == nil {
			http.Error(rw, "Body was nil", http.StatusBadRequest)
			return
		}
		var request model.Request
		// fmt.Printf("The body is: %v\n", r.Body)
		decode := json.NewDecoder(r.Body)
		// fmt.Println(decode)
		err := decode.Decode(&request)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error with request decoding %v\n", err), http.StatusBadRequest)
			return
		}
		// fmt.Println(request)

		rw.WriteHeader(http.StatusAccepted)
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SetupRoutes function
func SetupRoutes(apiPath string) {
	http.HandleFunc(fmt.Sprintf("%s/health-check", apiPath), healthCheckHandler)
	http.HandleFunc(fmt.Sprintf("%s/findGame", apiPath), findGameHandler)
}
