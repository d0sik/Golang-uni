package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"practice_2/internal/models"
)

var (
	tasks  = make(map[int]models.Task)
	nextID = 1
	mu     sync.Mutex
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
			return
		}

		task, ok := tasks[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
			return
		}

		json.NewEncoder(w).Encode(task)
		return
	}

	list := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}
	json.NewEncoder(w).Encode(list)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid title"})
		return
	}

	mu.Lock()
	task := models.Task{
		ID:    nextID,
		Title: input.Title,
		Done:  false,
	}
	tasks[nextID] = task
	nextID++
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	var input struct {
		Done bool `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mu.Lock()
	task, ok := tasks[id]
	if !ok {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}

	task.Done = input.Done
	tasks[id] = task
	mu.Unlock()

	json.NewEncoder(w).Encode(map[string]bool{"updated": true})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
		return
	}

	delete(tasks, id)
	json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
}
