package main

import (
	"math/rand"
	"net/http"
	"time"
)

func HttpGetMethodCheck(writer http.ResponseWriter, request *http.Request) bool {
	return request.Method == http.MethodGet
}

func HttpPostMethodCheck(writer http.ResponseWriter, request *http.Request) bool {
	return request.Method == http.MethodPost
}

func HttpPutMethodCheck(writer http.ResponseWriter, request *http.Request) bool {
	return request.Method == http.MethodPut
}

func HttpDeleteMethodCheck(writer http.ResponseWriter, request *http.Request) bool {
	return request.Method == http.MethodDelete
}

func GenerateId() int {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source).Intn(1000000)
}

func IsBookExist(id int) bool {
	for _, book := range books {
		if book.Id == id {
			return true
		}
	}
	return false
}
