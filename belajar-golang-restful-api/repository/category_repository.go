package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, db *sql.DB, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.DB, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.DB, categoryId int)
	FindById(ctx context.Context, tx *sql.DB, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.DB) []domain.Category
}