package repository

import (
	"context"
	"driver-pattern-repository/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment , error)
	FindById(ctx context.Context, id int32) (entity.Comment , error)
	FindAll(ctx context.Context) ([]entity.Comment , error)
}