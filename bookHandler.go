package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// making book variable to store the payload data
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

	// making default values
	book.Finished = book.PageCount == book.ReadPage
	book.Id = GenerateId()

	// push created book into array
	books = append(books, book)

	//validate if sucess
	_, isInserted := FindBook(books, book.Id)
	if !isInserted {
		http.Error(writer, "Failed to insert a new book, please ty again", http.StatusBadGateway)
		return
	}

	// http status, if created
	writer.WriteHeader(http.StatusCreated)
	// mapping custom response
	response := map[string]interface{}{
		"message": "book created",
		"data":    book,
	}

	// response in json
	json.NewEncoder(writer).Encode(response)
}

func ShowAllBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// check if array have values
	if len(books) == 0 {
		http.Error(writer, "There is no data yet", http.StatusOK)
		return
	}

	writer.WriteHeader(http.StatusOK)
	reponse := map[string]interface{}{
		"message": "Showing all books",
		"data":    books,
	}

	json.NewEncoder(writer).Encode(reponse)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	isGet := HttpGetMethodCheck(writer, request)
	if !isGet {
		http.Error(writer, request.Method+" Methid is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// getting params
	params := mux.Vars(request)
	bookId, paramsErr := strconv.Atoi(params["id"])
	if paramsErr != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	// finding book by its id
	bookData, isExist := FindBook(books, bookId)
	if !isExist {
		http.Error(writer, "id is not exist", http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "menampilkan buku dengan id",
		"data":    books[bookData],
	}

	json.NewEncoder(writer).Encode(response)
}

func EditBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	bookId, paramsErr := strconv.Atoi(params["id"])
	if paramsErr != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	bookData, isExist := FindBook(books, bookId)
	if !isExist {
		http.Error(writer, "id is not exist", http.StatusNotFound)
		return
	}

	var book Book

	bodyErr := json.NewDecoder(request.Body).Decode(&book)
	if bodyErr != nil {
		http.Error(writer, bodyErr.Error(), http.StatusBadRequest)
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

	// edit book data logic
	books[bookData].Title = book.Title
	books[bookData].Author = book.Author
	books[bookData].Publisher = book.Publisher
	books[bookData].PageCount = book.PageCount
	books[bookData].ReadPage = book.ReadPage
	books[bookData].Finished = book.PageCount == book.ReadPage

	writer.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "menampilkan data buku",
		"data":    books[bookData],
	}

	json.NewEncoder(writer).Encode(response)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	bookId, paramsErr := strconv.Atoi(params["id"])
	if paramsErr != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	bookIndex, isExist := FindBook(books, bookId)
	if !isExist {
		http.Error(writer, "id is not exist", http.StatusNotFound)
	}

	books = append(books[:bookIndex], books[bookIndex+1:]...)

	_, deleteFail := FindBook(books, bookId)
	if deleteFail {
		http.Error(writer, "deleting process failed", http.StatusBadGateway)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "book berhasil dihapus",
	}

	json.NewEncoder(writer).Encode(response)
}
