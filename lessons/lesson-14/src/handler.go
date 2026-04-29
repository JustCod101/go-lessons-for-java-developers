package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	repo *TaskRepository
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 简单的路由逻辑
	path := r.URL.Path
	if path == "/tasks" || path == "/tasks/" {
		switch r.Method {
		case http.MethodGet:
			h.listTasks(w, r)
		case http.MethodPost:
			h.createTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.HasPrefix(path, "/tasks/") {
		idStr := strings.TrimPrefix(path, "/tasks/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.getTask(w, r, id)
		case http.MethodPut:
			h.updateTask(w, r, id)
		case http.MethodDelete:
			h.deleteTask(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.repo.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := h.repo.Create(input.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, id int) {
	task, err := h.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var input struct {
		Completed bool `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.repo.Update(id, input.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
