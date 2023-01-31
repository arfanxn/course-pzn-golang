package services

import (
	"context"
	"database/sql"
	"golang-dependency-injection/exceptions"
	"golang-dependency-injection/helpers"
	"golang-dependency-injection/models/apis"
	"golang-dependency-injection/models/domains"
	"golang-dependency-injection/repositories"

	"github.com/go-playground/validator"
)

/* Contract */
type CategoryServiceContract interface {
	Get(ctx context.Context) []apis.CategoryResponse
	Find(ctx context.Context, id int32) (apis.CategoryResponse, error)
	Save(ctx context.Context, request apis.CategoryCreateRequest) apis.CategoryResponse
	Update(ctx context.Context, request apis.CategoryUpdateRequest) apis.CategoryResponse
	Delete(ctx context.Context, id int32) bool
}

/* Implemetation */
type CategoryService struct {
	CategoryRepository repositories.CategoryRepositoryContract
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(repository repositories.CategoryRepositoryContract, db *sql.DB, validate *validator.Validate) CategoryServiceContract {
	return &CategoryService{
		CategoryRepository: repository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryService) Get(ctx context.Context) []apis.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	categories := service.CategoryRepository.Get(ctx, tx)

	var categoriesReponse []apis.CategoryResponse
	for _, category := range categories {
		categoriesReponse = append(
			categoriesReponse, apis.CategoryResponse{
				Id:   int32(category.Id),
				Name: category.Name,
			})
	}

	return categoriesReponse
}

func (service *CategoryService) Find(ctx context.Context, id int32) (apis.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, id)
	var categoryResponse apis.CategoryResponse
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	categoryResponse = apis.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return categoryResponse, nil
}

func (service *CategoryService) Save(ctx context.Context, request apis.CategoryCreateRequest) apis.CategoryResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category := domains.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return apis.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (service *CategoryService) Update(ctx context.Context, request apis.CategoryUpdateRequest) apis.CategoryResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	category = domains.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return apis.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (service *CategoryService) Delete(ctx context.Context, id int32) bool {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return service.CategoryRepository.Delete(ctx, tx, category) == nil // it returns nil if no error
}
