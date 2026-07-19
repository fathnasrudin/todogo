package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Task struct {
	Title string `json:"title"`
}

type CreateTaskInput struct {
	Title string `json:"title"`
}

type CreateTaskResponse struct {
	Message string `json:"message"`
}

var tasks []Task

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusCreated)

	var taskInput CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	tasks = append(tasks, Task(taskInput))

	if err := json.NewEncoder(w).Encode(CreateTaskResponse{Message: "Success create task item"}); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
}

func main() {
	// add initial tasks
	tasks = append(tasks, Task{Title: "Implement Get tasks"}, Task{Title: "Implement Create a task"})

	http.HandleFunc("GET /api/tasks", GetTasksHandler)
	http.HandleFunc("POST /api/tasks", createTasksHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
