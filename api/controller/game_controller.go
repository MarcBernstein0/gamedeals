package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/middleware"
	"github.com/MarcBernstein0/gamedeals/api/model"
)

// healthCheckHandler function
func healthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, `{"alive":true}`)
}

func findGameHandler(rw http.ResponseWriter, r *http.Request) {
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
			log.Printf("%T, %v\n", err, err)
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
	healthHandler := http.HandlerFunc(healthCheckHandler)
	http.Handle(fmt.Sprintf("%s/health-check", apiPath), middleware.HeaderMiddleWare(healthHandler))
}
