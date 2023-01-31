package services

import (
	"errors"
	"golang-module/entities"
	"golang-module/repositories"
)

type CategoryService struct {
	Repository repositories.RepositoryInterface
}

func (service CategoryService) Get(id string) (*entities.CategoryEntity, error) {
	category := service.Repository.FindById(id)
	if category != nil {
		return category, nil
	} else {
		return nil, errors.New("Category not found")
	}

}
