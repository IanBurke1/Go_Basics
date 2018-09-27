package main

import (
	"fmt"
	"net/http"
)

// Index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World<h1>")
}

// about page
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About<h1>")
}

func main() {
	// Routing
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server Starting...")
	//Run web server on port 3000
	http.ListenAndServe(":3000", nil)

	// When running, Open web browser and enter localhost:3000
}
