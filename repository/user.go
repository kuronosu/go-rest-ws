package repository

import (
	"context"

	"github.com/kuronosu/go-rest-ws/models"
)

type UserRepository interface {
	ModelRepository[models.User]
}

var implementation UserRepository

func SetUserRepository(repo UserRepository) {
	implementation = repo
}

func InsertUser(ctx context.Context, data *models.User) error {
	return implementation.Insert(ctx, data)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
