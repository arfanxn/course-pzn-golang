package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database/models"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
 * Get database
 */
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

func TestCustomerRepositoryCreate(t *testing.T) {
	ctx := context.Background()
	customer := models.Customer{
		Name:      "Sparrow",
		Email:     sql.NullString{"sparrow@gm.com", true},
		Balance:   1000.00,
		Rating:    sql.NullFloat64{5.00, true},
		BirthDate: sql.NullTime{time.Now(), true},
		IsMarried: false,
		CreatedAt: sql.NullTime{time.Now(), true},
	}
	db := GetDB()
	repository := NewCustomerRepository(db)
	result, err := repository.Create(ctx, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCustomerRepositoryFind(t *testing.T) {
	ctx := context.Background()
	db := GetDB()

	repository := NewCustomerRepository(db)

	customer, err := repository.Find(ctx, int32(1))
	if err != nil {
		panic(err)
	}

	fmt.Println("Customer")
	fmt.Println("Customer ID:", customer.Id)
}

func TestCustomerRepositoryGet(t *testing.T) {
	ctx := context.Background()
	db := GetDB()

	repository := NewCustomerRepository(db)

	customers, err := repository.Get(ctx)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(customers); i++ {
		fmt.Println("Customer ID:", customers[i].Id)
	}
}
