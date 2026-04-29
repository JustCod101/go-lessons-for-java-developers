package main

import (
	"fmt"
	"net/http"
)

func main() {
	repo := NewTaskRepository()
	handler := &TaskHandler{repo: repo}

	fmt.Println("Task API starting at :8080...")
	http.ListenAndServe(":8080", handler)
}
