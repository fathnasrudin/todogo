package todo

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// write headers
	w.Header().Set("Content-Type", "application/json")

	// should validate id

	// delete task
	if err := deleteTask(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(BadResponse{Message: err.Error()})
		return
	}

	// delete task
	json.NewEncoder(w).Encode(UpdateTaskResponse{Message: "Success delete task: " + id})
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(Tasks)

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

	Tasks = append(Tasks, NewTask(taskInput))

	if err := json.NewEncoder(w).Encode(CreateTaskResponse{Message: "Success create task item"}); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
}
