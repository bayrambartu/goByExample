package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	/*
		r := chi.NewRouter()

		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello WORLD !"))
		})

		http.ListenAndServe(":3000", r) */

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home page- Welcome!"))
	})
	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello " + name,
		})
	})
	http.ListenAndServe(":3000", r)
}
