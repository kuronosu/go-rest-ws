package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/kuronosu/go-rest-ws/errors"
	"github.com/kuronosu/go-rest-ws/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func (repo *UserPostgresRepository) Insert(ctx context.Context, data *models.User) error {
	row := repo.db.QueryRowContext(ctx,
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		data.Email, data.Password)
	err := row.Scan(&data.Id)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return errors.UserAlreadyExistError
		}
	}
	return err
}

func (repo *UserPostgresRepository) GetById(ctx context.Context, id int64) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM TABLE users WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err) // TODO better handle error
		}
	}()
	if err != nil {
		return nil, err
	}
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserPostgresRepository) Close() error {
	return repo.db.Close()
}

type PostgresRepository struct {
	User *UserPostgresRepository
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		User: &UserPostgresRepository{db},
	}, nil
}
