package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	GetDB()
}

func TestDBExecSql(t *testing.T) {
	db := GetDB()
	defer db.Close()

	ctx := context.Background()

	result, err := db.ExecContext(
		ctx,
		"INSERT INTO customers(name, email, balance, rating, birth_date, is_married, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		"dorsey",
		"dorsey@gmail.com",
		1000.00,
		5.00,
		"2020-10-10",
		false,
		"2022-10-10",
	)

	if err != nil {
		panic("error inserting customer")
	}

	fmt.Println("successfully inserting customer")

	// Get last inserted id
	lastInsertedId, _ := result.LastInsertId()
	fmt.Println("Last inserted Customer's id:", lastInsertedId)

}

func TestDBQuerySql(t *testing.T) {
	db := GetDB()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(
		ctx,
		"SELECT id, name, email, balance, rating, birth_date, is_married, created_at FROM customers")
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance float64
		var rating float32
		var isMarried bool
		var birthDate sql.NullTime
		var createdAt time.Time
		err := rows.Scan(
			&id, &name, &email, &balance, &rating, &birthDate, &isMarried, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("Customer")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth date:", birthDate)
		fmt.Println("Is married:", isMarried)
		fmt.Println("Created at:", createdAt)
	}

}

func TestDBPrepStmt(t *testing.T) {
	db := GetDB()
	defer db.Close()

	ctx := context.Background()

	query :=
		"INSERT INTO customers(name, email, balance, rating, birth_date, is_married, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		unixPoachAsString := strconv.FormatInt(time.Now().Unix(), 10)
		result, err := stmt.ExecContext(ctx,
			"dorsey"+unixPoachAsString,
			"dorsey"+unixPoachAsString+"@gmail.com",
			1000.00,
			5.00,
			"2020-10-10",
			false,
			"2022-10-10")
		if err != nil {
			panic(err)
		}

		lastInsertedId, _ := result.LastInsertId()
		fmt.Println("Last inserted customer's id:", lastInsertedId)

		time.Sleep(time.Second) //  a second
	}

}

func TestDBTransaction(t *testing.T) {
	db := GetDB()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO customers(name, email, balance, rating, birth_date, is_married, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		"alpine",
		"alpine@gmail.com",
		1000.00,
		5.00,
		"2020-10-10",
		false,
		"2022-10-10",
	)
	if err != nil {
		panic(err)
	}

	if tx.Commit() != nil {
		tx.Rollback()
	}

	// Get last inserted id
	lastInsertedId, _ := result.LastInsertId()
	fmt.Println("Last inserted Customer's id:", lastInsertedId)

}
