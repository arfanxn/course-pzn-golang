package repositories

import (
	"golang-module/entities"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entities.CategoryEntity {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) != nil {
		category := arguments.Get(0).(entities.CategoryEntity)
		return &category
	} else {
		return nil
	}
}
