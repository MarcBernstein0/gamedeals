package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarcBernstein0/gamedeals/api/controller"
)

func main() {
	http.HandleFunc("api/test_connection", controller.TestHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Server failed\n%v\n", err)
	}
	fmt.Printf("Hello world\n")
}
