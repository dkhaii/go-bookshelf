package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	PageCount int    `json:"pageCount"`
	ReadPage  int    `json:"readPage"`
	Reading   bool   `json:"reading"`
}

var books []Book

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome gerin")
}

func insert(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method != http.MethodPost {
		http.Error(writer, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	var book Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	books = append(books, book)
	writer.WriteHeader(http.StatusCreated)
	
	err2 := json.NewEncoder(writer).Encode(books)
	if err2 != nil {
		http.Error(writer, err2.Error(), http.StatusBadGateway)
	}
}

func main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/insert", insert)

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
