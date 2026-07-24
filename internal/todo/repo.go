package todo

import (
	"database/sql"
	"errors"
	"log"
)

type TodoRepository interface {
	List() ([]Task, error)
	Create(Task) error
	Update(id string, data UpdateTaskInput) error
	Delete(id string ) error
}

type PostgresTodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *PostgresTodoRepository{
	return &PostgresTodoRepository{
		db: db,
	}
}

func (r *PostgresTodoRepository) List() ([]Task, error) {
	var task Task
	var tasks []Task

	// call 
	rows, queryErr := r.db.Query(`
	SELECT id, title, is_done
	FROM todos
	`);

	if queryErr != nil {
		log.Println("Failed to run query", queryErr)
	}

	defer rows.Close()

	// convert db row into Task struct in tasks slice
	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.IsDone)
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

func (r *PostgresTodoRepository)  Create( task Task) error {
	query := `INSERT INTO todos (id, title) VALUES ($1, $2);`;
	_, err := r.db.Exec(query, task.ID, task.Title )
	if err != nil {return err}

	return nil
}

func (r *PostgresTodoRepository)  Update(id string, task UpdateTaskInput) error {
	query := `
		UPDATE todos
		SET 
			title = COALESCE($1, title),
			is_done = COALESCE($2, is_done)
		WHERE 
			id = $3;
	`;
	_, err := r.db.Exec(query, task.Title, task.IsDone, id)
	if err != nil {return err}

	return nil
}

func (r *PostgresTodoRepository)  Delete(id string) error {
	query := `
		DELETE FROM todos
		WHERE id = $1;
	`;

	result, err := r.db.Exec(query, id);
	if err != nil {
		return err;
	}

	rowsAffected, err := result.RowsAffected();
	if err != nil {return err}
	if rowsAffected == 0 {
		return errors.New("Failed to delete task")
	}

	return nil
}