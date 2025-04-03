package main

import (
	"fmt"
	"net/http"
)

/*
	func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)

}
*/

/*
	func homeHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Home Page ")
	}

	func aboutHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is About Page ")
	}
*/

/*
	func queryHandler(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guests"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	}
*/
/*
type Response struct {
	Message string `json:"message"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{"Hello, this is a JSON response"}
	json.NewEncoder(w).Encode(resp)
}
*/

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello golang")
}
func header(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%s: %s\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", header)
	http.ListenAndServe(":8090", nil)

	/*
		http.HandleFunc("/json", jsonHandler)
		http.ListenAndServe(":8080", nil) */
	/*
		http.HandleFunc("/greet", queryHandler)

		fmt.Println("server is running")
		http.ListenAndServe(":8080", nil)
	*/

	/*
		http.HandleFunc("/", homeHandler)
		http.HandleFunc("/about", aboutHandler)

		fmt.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)*/
	/*
		http.HandleFunc("/", handler)
		fmt.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)
	*/
}
