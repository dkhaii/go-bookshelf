package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"

	routes := SetupRoutes()

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
