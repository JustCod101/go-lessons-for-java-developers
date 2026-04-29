package main

import (
	"errors"
	"sync"
)

type TaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (r *TaskRepository) Create(title string) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := Task{
		ID:        r.nextID,
		Title:     title,
		Completed: false,
	}
	r.tasks[task.ID] = task
	r.nextID++
	return task
}

func (r *TaskRepository) List() []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		list = append(list, t)
	}
	return list
}

func (r *TaskRepository) Get(id int) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *TaskRepository) Update(id int, completed bool) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}
	task.Completed = completed
	r.tasks[id] = task
	return task, nil
}

func (r *TaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
