package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{ID: 1, Name: "Java Developer"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
