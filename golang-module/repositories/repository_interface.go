package repositories

import (
	"golang-module/entities"
)

type RepositoryInterface interface {
	FindById(id string) *entities.CategoryEntity
}
