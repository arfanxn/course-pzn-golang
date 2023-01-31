package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function")
}

func GetDB() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"root:@tcp(localhost:3306)/golang_database?parseTime=true")

	if err != nil {
		fmt.Println("error opening database")
		// panic("database connection error: ")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
