package todo

import (
	"net/http"
)


func RegisterRoutes(mux *http.ServeMux) {
	// db, dbErr := database.Open()
	// if dbErr != nil {
	// 	log.Fatal(dbErr)
	// }
	// defer db.Close()

	// repo := NewTodoRepository(db)
	// service := NewTaskService(repo)
	// handler := NewTodoHandler(*service)
	// mux.HandleFunc("GET /api/tasks", handler.List)
	// mux.HandleFunc("POST /api/tasks", handler.Create)
	// mux.HandleFunc("PUT /api/tasks/{id}", handler.Update)
	// mux.HandleFunc("DELETE /api/tasks/{id}", handler.Delete)
}