package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBconnection(username, password, localhost, port, dbname string) *sql.DB {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, localhost, port, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}

	return db
}
