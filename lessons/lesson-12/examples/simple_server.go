package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", nil)
}
