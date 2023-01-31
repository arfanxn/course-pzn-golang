package app

import (
	"database/sql"
	"golang-dependency-injection/helpers"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	/*
	* Run this for database migration
	* migrate -database "mysql://root@tcp(localhost:3306)/golang_database" -path databases/migrations up
	 */

	return db
}

func NewTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	return db
}
