package repository

import (
	"app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUserById(id int) (app.User, error)
	PullOutSessionByGUID(fingerprint string) (app.Sesion, error)
	CreateSession(session app.Sesion) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
