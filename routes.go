package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome gerin")
}

func SetupRoutes() *mux.Router {
	routes := mux.NewRouter()

	subRoutes := routes.PathPrefix("/api/v1").Subrouter()
	subRoutes.HandleFunc("/", index).Methods("GET")

	bookRoutes := subRoutes.PathPrefix("/book").Subrouter()
	bookRoutes.HandleFunc("/insert", InsertBook).Methods("POST")
	bookRoutes.HandleFunc("/show-all", ShowAllBooks).Methods("GET")
	bookRoutes.HandleFunc("/find/{id}", GetBookById).Methods("GET")
	bookRoutes.HandleFunc("/edit/{id}", EditBook).Methods("PUT")
	bookRoutes.HandleFunc("/delete/{id}", DeleteBook).Methods("DELETE")

	return routes
}
