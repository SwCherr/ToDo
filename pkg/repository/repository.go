package repository

import (
	todo "app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error) // return id new user in db & error
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItems interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
