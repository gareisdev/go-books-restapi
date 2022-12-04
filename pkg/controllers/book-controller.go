package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gareisdev/go-books-restapi/pkg/models"
	"github.com/gareisdev/go-books-restapi/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllBooks(response http.ResponseWriter, request *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	book_id := vars["id"]

	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		log.Fatal("[ERROR] Invalid ID")
	}

	book, _ := models.GetBook(ID)
	res, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	newBook := &models.Book{}

	utils.ParseBody(request, newBook)
	book := newBook.CreateBook()

	res, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	book_id := vars["id"]

	ID, err := strconv.ParseInt(book_id, 0, 0)

	if err != nil {
		log.Fatal("[ERROR] Invalid ID")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {

	var updateBook = &models.Book{}

	utils.ParseBody(request, updateBook)

	vars := mux.Vars(request)
	book_id := vars["id"]

	ID, err := strconv.ParseInt(book_id, 0, 0)

	if err != nil {
		log.Fatal("[ERROR] Invalid ID")
	}

	book, db := models.GetBook(ID)

	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
