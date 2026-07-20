package main

import (
	"log"
	"net/http"

	"github.com/fathnasrudin/todogo/internal/todo"
)

func main() {
	mux := http.NewServeMux();

	todo.RegisterRoutes(mux);

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
