package routes

import (
	"github.com/gareisdev/go-books-restapi/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBooksRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
