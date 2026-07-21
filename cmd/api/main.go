package main

import (
	"log"
	"net/http"

	"github.com/fathnasrudin/todogo/internal/common/database"
	"github.com/fathnasrudin/todogo/internal/todo"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, dbErr := database.Open()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	mux := http.NewServeMux();

	repo := todo.NewTodoRepository(db)
	service := todo.NewTaskService(repo)
	handler := todo.NewTodoHandler(*service)
	mux.HandleFunc("GET /api/tasks", handler.List)
	mux.HandleFunc("POST /api/tasks", handler.Create)
	mux.HandleFunc("PUT /api/tasks/{id}", handler.Update)
	mux.HandleFunc("DELETE /api/tasks/{id}", handler.Delete)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
