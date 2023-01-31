package services

import (
	"golang-module/entities"
	"golang-module/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	t.Run("ShouldReturnNil", func(t *testing.T) {
		categoryRepository.Mock.On("FindById", "1").Return(nil)
		category, err := categoryService.Get("1")
		assert.Nil(t, category)
		assert.NotNil(t, err)
	})

	t.Run("ShouldReturnData", func(t *testing.T) {
		category := entities.CategoryEntity{
			Id:   "2",
			Name: "Electronic",
		}

		categoryRepository.Mock.On("FindById", "2").Return(category)
		returnedCategory, err := categoryService.Get("2")

		assert.NotNil(t, category)
		assert.Equal(t, category.Id, returnedCategory.Id)
		assert.Nil(t, err)
	})
}
