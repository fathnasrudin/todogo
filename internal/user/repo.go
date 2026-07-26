package user

import (
	"database/sql"
	"errors"
)

type UserRepo interface {
	List() ([]GetUserResponse, error)
	Create(u User) (error)
	Update(id string, u UpdateUserInput) (error)
	Delete(id string) (error)
}

type PostgresUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

func (r *PostgresUserRepo) List() ([]GetUserResponse, error) {
	var userData GetUserResponse;
	var users []GetUserResponse

	// takedata from db
	query := `
	SELECT id, email
	FROM users;
	`

	rows, err := r.db.Query(query);
	if err != nil {return nil, err}

	for rows.Next() {
		err := rows.Scan(&userData.ID, &userData.Email)
		if err != nil {
			{return nil, err}
		}
		users = append(users, userData)
	}
	
	return users, nil
}


func (r *PostgresUserRepo) Create(u User) ( error) {
	query := `
	INSERT INTO users (id, email, password)
	VALUES ($1, $2, $3);
	`
	_, err := r.db.Exec(query, u.ID, u.Email, u.Password);
	if err != nil {return err}
	
	return nil
}

func (r *PostgresUserRepo) Update(id string, u UpdateUserInput) ( error) {
	query := `
	UPDATE users
	SET 
		email = COALESCE($1, email),
		password = COALESCE($2, password)
	WHERE id = $3;
	`
	result, err := r.db.Exec(query, u.Email, u.Password, id);

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Task not found")
	}

	return err
}

func (r *PostgresUserRepo) Delete(id string) ( error) {
	query := `
	DELETE FROM users
	WHERE id = $1;
	`
	result, err := r.db.Exec(query, id);

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Task not found")
	}

	return err
}