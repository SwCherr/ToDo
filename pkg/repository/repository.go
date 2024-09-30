package repository

import (
	todo "app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error) // return id new user in db & error
	GetUser(username, password string) (todo.User, error)
	GetUserById(id int) (todo.User, error)
	CreateSession(user_id int, user_ip, token string) error
	DeleteSession(user todo.User) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
