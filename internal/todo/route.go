package todo

import "net/http"


func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/tasks", GetTasksHandler)
	mux.HandleFunc("POST /api/tasks", CreateTasksHandler)
	mux.HandleFunc("PUT /api/tasks/{id}", UpdateTaskHandler)
	mux.HandleFunc("DELETE /api/tasks/{id}", DeleteTaskHandler)
}