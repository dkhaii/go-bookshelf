package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	PageCount int    `json:"pageCount"`
	ReadPage  int    `json:"readPage"`
	Finished  bool   `json:"Finished"`
}

var books []Book

func InsertBook(writer http.ResponseWriter, request *http.Request) {
	// setting the header to json
	writer.Header().Set("Content-Type", "application/json")

	// method checking
	isPost := HttpPostMethodCheck(writer, request)
	if !isPost {
		http.Error(writer, request.Method+" Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// making book variable to store the created data into array of books
	var book Book

	// request body
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// validation
	if book.Title == "" {
		http.Error(writer, "Judul is required", http.StatusBadRequest)
		return
	}

	if book.ReadPage > book.PageCount {
		http.Error(writer, "readPage must not be above pageCount", http.StatusBadRequest)
		return
	}

	book.Finished = book.PageCount == book.ReadPage
	book.Id = GenerateId()

	// push created book into array
	books = append(books, book)
	//validate if success
	isInserted := IsBookExist(book.Id)
	if !isInserted {
		http.Error(writer, "Failed to insert a new book", http.StatusBadGateway)
		return
	}

	// http status, if created
	writer.WriteHeader(http.StatusCreated)
	// mapping the response
	response := map[string]interface{}{
		"message": "book created",
		"data":    book,
	}

	// response in json
	json.NewEncoder(writer).Encode(response)
}

func ShowAllBooks(writer http.ResponseWriter, request *http.Request) {
	// setting the header to json
	writer.Header().Set("Content-Type", "application/json")

	// method checking
	isGet := HttpGetMethodCheck(writer, request)
	if !isGet {
		http.Error(writer, request.Method+" Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// check if array have values
	if len(books) == 0 {
		http.Error(writer, "There is no data yet", http.StatusOK)
		return
	}

	// http status
	writer.WriteHeader(http.StatusOK)
	// mapping the response
	reponse := map[string]interface{}{
		"message": "Showing all books",
		"data":    books,
	}

	// reponse in json
	json.NewEncoder(writer).Encode(reponse)
}
