package main

import (
	"log"
	"net/http"

	"github.com/gareisdev/go-books-restapi/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBooksRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
