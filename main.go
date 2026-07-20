package main

import (
	"log"
	"net/http"

	"github.com/fathnasrudin/todogo/todo"
)

func main() {

	// add initial tasks
	todo.Tasks = append(todo.Tasks,
		(todo.Task{Title: "Implement Create task", ID: "019f7b2a-d971-7227-b295-e7088449e296"}),
		(todo.Task{Title: "Implement Get tasks", ID: "019f7b2a-d971-722c-b0e6-14d1aa6bf334"}))

	http.HandleFunc("GET /api/tasks", todo.GetTasksHandler)
	http.HandleFunc("POST /api/tasks", todo.CreateTasksHandler)
	http.HandleFunc("PUT /api/tasks/{id}", todo.UpdateTaskHandler)
	http.HandleFunc("DELETE /api/tasks/{id}", todo.DeleteTaskHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
