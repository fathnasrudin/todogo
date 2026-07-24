package todo

import (
	"encoding/json"
	"log"
	"net/http"
)

type ITodoHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	service TaskService
}

func NewTodoHandler(service TaskService) *TodoHandler{
	return &TodoHandler{
		service: service,
	}
}


func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	if err := h.service.Update(id, updateTaskInput); err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(UpdateTaskResponse{Message: err.Error()})
		return
	}

	// success update
	json.NewEncoder(w).Encode(UpdateTaskResponse{Message: "Success update task"})
}

func (h *TodoHandler)  Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// write headers
	w.Header().Set("Content-Type", "application/json")

	// should validate id

	// delete task
	if err := h.service.Delete(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(BadResponse{Message: err.Error()})
		return
	}

	// delete task
	json.NewEncoder(w).Encode(UpdateTaskResponse{Message: "Success delete task: " + id})
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// service
	tasks, tasksErr := h.service.List()
	if tasksErr != nil {
		http.Error(w, tasksErr.Error(), http.StatusInternalServerError)
	}

	err := json.NewEncoder(w).Encode(tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler)  Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var taskInput CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Call service
	if err := h.service.Create(taskInput); err != nil {
		http.Error(w, "Failed to Create Task" + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(CreateTaskResponse{Message: "Success create task item"}); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
}
