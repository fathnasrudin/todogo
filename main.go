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
	Title string
}

var tasks []Task

func main() {

	// add tasks
	tasks = append(tasks, Task{Title: "Implement Get tasks"},  Task{Title: "Implement Create a task"}, )

	http.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(tasks)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
