package service

import (
	"golang-unit-testing/entity"
	"golang-unit-testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	
	// Start unit test
	result, err := categoryService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	
}

func TestCategoryService_GetSuccess(t *testing.T) {
	// Create category
	category := entity.Category {
		Id: "2",
		Name: "Laptop",
	}

	// Program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// Start unit test
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}