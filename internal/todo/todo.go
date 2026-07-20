package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type BadResponse struct {
	Message string `json:"message"`
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

var Tasks []Task

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

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(Tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {
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

func findTask(taskId string) (*Task, error) {
	var foundTask *Task

	// find task based on id
	for i := range Tasks {
		if Tasks[i].ID == taskId {
			foundTask = &Tasks[i]
			break
		}
	}

	if foundTask == nil {
		return nil, fmt.Errorf("Task with ID %q not found", taskId)
	}

	return foundTask, nil
}

func deleteTask(taskId string) error {
	// find task
	task, err := findTask(taskId)
	if err != nil {
		return err
	}

	// delete task
	for i := range Tasks {
		if Tasks[i].ID == task.ID {
			Tasks = slices.Delete(Tasks, i, i+1)
			break
		}
	}

	return nil
}

func updateTask(taskId string, tData UpdateTaskInput) error {
	var foundTask *Task

	// find task based on id
	for i := range Tasks {
		if Tasks[i].ID == taskId {
			foundTask = &Tasks[i]
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

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
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

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
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

func init() {
	// add initial tasks
	Tasks = append(Tasks,
		(Task{Title: "Implement Create task", ID: "019f7b2a-d971-7227-b295-e7088449e296"}),
		(Task{Title: "Implement Get tasks", ID: "019f7b2a-d971-722c-b0e6-14d1aa6bf334"}))
}