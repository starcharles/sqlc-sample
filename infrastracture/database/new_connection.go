package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	return db
}
