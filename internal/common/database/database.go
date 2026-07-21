package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

	fmt.Println("Success connect to db")

    return db, nil
}