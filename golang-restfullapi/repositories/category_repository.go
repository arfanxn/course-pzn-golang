package repositories

import (
	"context"
	"database/sql"
	"golang-restfullapi/helpers"
	"golang-restfullapi/models/domains"
)

/* Contract */
type CategoryRepositoryContract interface {
	Get(ctx context.Context, tx *sql.Tx) []domains.Category
	Find(ctx context.Context, tx *sql.Tx, id int32) (domains.Category, error)
	Save(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category
	Update(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category
	Delete(ctx context.Context, tx *sql.Tx, category domains.Category) error
}

/* Implemetation */
type CategoryRepository struct {
	//
}

func NewCategoryRepository() CategoryRepositoryContract {
	return &CategoryRepository{}
}

func (repository *CategoryRepository) Get(ctx context.Context, tx *sql.Tx) []domains.Category {
	query := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, query)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories []domains.Category
	for rows.Next() {
		category := domains.Category{}
		_ = rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	return categories
}

func (repository *CategoryRepository) Find(ctx context.Context, tx *sql.Tx, id int32) (domains.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	defer rows.Close()

	category := domains.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, err
	}
}

func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category {
	query := "INSERT INTO categories(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	category.Id = int32(id)
	return category
}

func (repository *CategoryRepository) Update(ctx context.Context, tx *sql.Tx, category domains.Category) domains.Category {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helpers.PanicIfError(err)
	return category
}

func (repository *CategoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domains.Category) error {
	query := "DELETE FROM categories WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, category.Id)
	helpers.PanicIfError(err)

	affected, err := result.RowsAffected()

	if affected > 0 {
		return nil
	} else {
		return err
	}
}
