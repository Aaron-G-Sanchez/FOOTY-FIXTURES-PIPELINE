package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to db: ", err)
	}

	return db
}
