package repository

import (
	"golang-unit-testing/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (crm *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := crm.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
	}
}