package todo

import (
	"database/sql"
	"log"
)

type ITodoRepository interface {
	List() ([]Task, error)
	Create(Task) error
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository{
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) List() ([]Task, error) {
	var task Task
	var tasks []Task

	// call 
	rows, queryErr := r.db.Query(`
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
		tasks = append(tasks, task)
	}

	if queryErr != nil {
		return nil, queryErr;
	}
	return tasks, nil;
}

func (r *TodoRepository)  Create( task Task) error {
	query := `INSERT INTO todos (id, title) VALUES ($1, $2);`;
	_, err := r.db.Exec(query, task.ID, task.Title )
	if err != nil {return err}

	return nil
}