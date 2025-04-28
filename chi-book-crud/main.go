package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// In-memory book list
var books = []Book{
	{ID: "1", Title: "First Book", Author: "Steinbeck "},
	{ID: "2", Title: "Second Book", Author: "Steinbeck "},
}

func main() {
	r := chi.NewRouter()
	// Routes
	r.Get("/books", getBooks)
	r.Get("/books/{id}", getBookByID)
	r.Post("/books", createBook)
	r.Put("/books/{id}", updateBook)
	r.Delete("/books/{id}", deleteBook)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)

}

// handler's
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
func getBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedBook Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "book not found ", http.StatusNotFound)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "book not found", http.StatusNotFound)
}
