package main

import (
	"log"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/controller"
)

func main() {
	log.Println("Starting server")
	controller.SetupRoutes("/api")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Server failed\n%v\n", err)
	}
}
