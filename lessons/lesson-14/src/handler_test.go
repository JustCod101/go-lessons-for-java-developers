package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTask(t *testing.T) {
	repo := NewTaskRepository()
	handler := &TaskHandler{repo: repo}

	body := `{"title": "Learn Go"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "Learn Go") {
		t.Errorf("Response body does not contain task title")
	}
}
