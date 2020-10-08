package main

import (
	"log"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/controller"
	"github.com/MarcBernstein0/gamedeals/api/middleware"
)

func main() {
	log.Println("Starting server")
	healthHandler := http.HandlerFunc(controller.HealthCheckHandler)
	http.Handle("/api/health-check", middleware.HeaderMiddleWare(healthHandler))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Server failed\n%v\n", err)
	}
}
