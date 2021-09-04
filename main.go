package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model) = Is like a Class can have a property and methods
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get All Books
func getBooks() {

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandlerFunc("/api/books", getBooks).Methods("GET")
	r.HandlerFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandlerFunc("/api/books", createBook).Methods("POST")
	r.HandlerFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandlerFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
