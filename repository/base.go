package repository

import (
	"context"

	"github.com/kuronosu/go-rest-ws/models"
)

type ModelRepository[T models.Models] interface {
	Insert(ctx context.Context, data *T) error
	GetById(ctx context.Context, id int64) (*T, error)
	Close() error
}
