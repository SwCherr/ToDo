package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
