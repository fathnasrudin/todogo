package user

import "database/sql"

type UserRepo interface {
	List() ([]GetUserResponse, error)
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