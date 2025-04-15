package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
	{ID: "2", Title: "Sapiens", Author: "Yuval Noah Harari"},
}

func main() {

	http.HandleFunc("/books", getBooks)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
