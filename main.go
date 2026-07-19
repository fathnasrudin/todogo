package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Task struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

type CreateTaskInput struct {
	Title string `json:"title"`
}

type UpdateTaskInput struct {
	Title string `json:"title"`
}

type UpdateTaskResponse struct {
	Message string `json:"message"`
}

type CreateTaskResponse struct {
	Message string `json:"message"`
}

func NewTask(t CreateTaskInput) Task {
	IDByte, err := uuid.NewV7()

	if err != nil {
		log.Fatalln("failed to generate ID", err)
	}
	return Task{
		Title: t.Title,
		ID:    IDByte.String(),
	}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var taskInput CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	tasks = append(tasks, NewTask(taskInput))

	if err := json.NewEncoder(w).Encode(CreateTaskResponse{Message: "Success create task item"}); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
}

func updateTask(taskId string, tData UpdateTaskInput) error {
	var foundTask *Task

	// find task based on id
	for i := range tasks {
		if tasks[i].ID == taskId {
			foundTask = &tasks[i]
			break
		}
	}

	if foundTask == nil {
		return fmt.Errorf("Task with ID %q not found", taskId)
	}

	// task found
	// update task
	foundTask.Title = tData.Title

	// no error mean success update
	return nil
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")

	// should validate id

	// decode body input
	var updateTaskInput UpdateTaskInput
	err := json.NewDecoder(r.Body).Decode(&updateTaskInput)
	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(UpdateTaskResponse{Message: err.Error()})
		return
	}

	// update input
	if err := updateTask(id, updateTaskInput); err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(UpdateTaskResponse{Message: err.Error()})
		return
	}

	// success update
	json.NewEncoder(w).Encode(UpdateTaskResponse{Message: "Success update task"})
}

func main() {
	// add initial tasks
	tasks = append(tasks,
		NewTask(CreateTaskInput{Title: "Implement Create task"}),
		NewTask(CreateTaskInput{Title: "Implement Get tasks"}))

	http.HandleFunc("GET /api/tasks", GetTasksHandler)
	http.HandleFunc("POST /api/tasks", createTasksHandler)
	http.HandleFunc("PUT /api/tasks/{id}", updateTaskHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
