package service

import (
	"errors"
	"golang-unit-testing/entity"
	"golang-unit-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (cs CategoryService) Get(id string) (*entity.Category, error) {
	category := cs.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}