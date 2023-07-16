package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome gerin")
}

func main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1", index)
	mux.HandleFunc("/api/v1/insert", InsertBook)
	mux.HandleFunc("/api/v1/showAll", ShowAllBooks)

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
