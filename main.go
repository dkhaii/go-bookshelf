package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome gerin")
}

func main() {
	port := "8080"

	routes := mux.NewRouter()

	routes.HandleFunc("/api/v1", index)
	routes.HandleFunc("/api/v1/book/insert", InsertBook)
	routes.HandleFunc("/api/v1/book/show-all", ShowAllBooks)
	routes.HandleFunc("/api/v1/book/{id}", EditBook)

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
