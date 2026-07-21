package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fathnasrudin/todogo/internal/common/database"
	"github.com/fathnasrudin/todogo/internal/todo"
	_ "github.com/jackc/pgx/v5/stdlib"
)


func repoGetTodoList(db *sql.DB) error{
	var task todo.Task

	// call 
	rows, queryErr := db.Query(`
	SELECT *
	FROM todos
	`);

	if queryErr != nil {
		log.Println("Failed to run query", queryErr)
	}

	defer rows.Close()

	// convert db row into Task struct in tasks slice
	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title)
		if err != nil {
			log.Fatal(err);
		}

		// append tasks
		todo.Tasks = append(todo.Tasks, task)
	}

	if queryErr != nil {
		return queryErr;
	}
	return nil;
}

func main() {
	db, dbErr := database.Open()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	if err := repoGetTodoList(db); err != nil {
		log.Fatal("Failed to get list of todo: ", err)
	}

	mux := http.NewServeMux();

	todo.RegisterRoutes(mux);

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to load server: ", err)
	}
}
