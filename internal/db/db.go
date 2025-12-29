package db

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	connStr := "user=postgres password=postgres dbname=testdb sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
