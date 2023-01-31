package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/models"
	"strconv"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer models.Customer) (models.Customer, error)
	Find(ctx context.Context, id int32) (models.Customer, error)
	Get(ctx context.Context) ([]models.Customer, error)
}

type customerRepositoryImplementation struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImplementation{DB: db}
}

/*
 *	Implementation methods
 */

func (repository *customerRepositoryImplementation) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	query := "INSERT INTO customers(name, email, balance, rating, birth_date, is_married, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, query,
		customer.Name,
		customer.Email,
		customer.Balance,
		customer.Rating,
		customer.BirthDate,
		customer.IsMarried,
		customer.CreatedAt)
	if err != nil {
		return customer, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return customer, err
	}

	customer.Id = int32(id)

	return customer, nil
}

func (repository *customerRepositoryImplementation) Find(ctx context.Context, id int32) (models.Customer, error) {
	db := repository.DB
	rows, err := db.QueryContext(
		ctx,
		"SELECT id, name, email, balance, rating, birth_date, is_married, created_at FROM customers WHERE id=?", id)
	defer rows.Close()
	customer := models.Customer{}
	if err != nil {
		return customer, err
	}

	if rows.Next() {
		rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.Balance,
			&customer.Rating,
			&customer.BirthDate,
			&customer.IsMarried,
			&customer.CreatedAt)
		return customer, nil

	} else {
		return customer, errors.New("Customer with ID " + strconv.Itoa(int(id)) + " are not found")
	}
}

func (repository *customerRepositoryImplementation) Get(ctx context.Context) ([]models.Customer, error) {
	db := repository.DB
	rows, err := db.QueryContext(
		ctx,
		"SELECT id, name, email, balance, rating, birth_date, is_married, created_at FROM customers")
	defer rows.Close()
	var customers []models.Customer
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		customer := models.Customer{}
		rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.Balance,
			&customer.Rating,
			&customer.BirthDate,
			&customer.IsMarried,
			&customer.CreatedAt)
		customers = append(customers, customer)
	}

	return customers, nil
}

/*
 *	End of implementation methods
 */
