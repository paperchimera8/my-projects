package db

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=learnle sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
