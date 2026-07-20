package todo

import "net/http"


func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/tasks", getTasksHandler)
	mux.HandleFunc("POST /api/tasks", createTasksHandler)
	mux.HandleFunc("PUT /api/tasks/{id}", updateTaskHandler)
	mux.HandleFunc("DELETE /api/tasks/{id}", deleteTaskHandler)
}