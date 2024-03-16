package repository

import "mock-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}