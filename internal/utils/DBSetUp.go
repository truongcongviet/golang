package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 5432
	username      = "postgres"
	password      = "quanpro99"
	database_name = "friend_management"
)

func DBConnection() (db *sql.DB) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, database_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	return db
}
