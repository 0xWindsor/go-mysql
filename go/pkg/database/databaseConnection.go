package databaseConnection

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to the MySQL database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1q2w3E*@tcp(localhost:3306)/tasktally")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
