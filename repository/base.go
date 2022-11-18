package repository

import (
	"context"

	"github.com/kuronosu/go-rest-ws/models"
)

type ModelRepository[T models.Models[T]] interface {
	Insert(ctx context.Context, data *T) error
	GetById(ctx context.Context, id int64) (*T, error)
	Close() error
}

// type Repository interface {
// 	CrudRepository
// }
//
// type CrudRepository interface {
// 	Insert(ctx context.Context, data interface{}) error
// 	GetById(ctx context.Context, id int64) (*Model, error)
// 	Close() error
// }
