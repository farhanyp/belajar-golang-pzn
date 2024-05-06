package repository

import (
	"belajar-golang-restful-api/model/web"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, db *sql.DB, category web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, db *sql.DB, category web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, db *sql.DB, categoryId int)
	FindById(ctx context.Context, db *sql.DB, categoryId int) (web.CategoryResponse, error)
	FindAll(ctx context.Context, db *sql.DB) []web.CategoryResponse
}