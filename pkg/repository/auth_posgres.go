package repository

import (
	todo "app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateSession(user todo.User) error {
	var id int
	query := fmt.Sprintf("UPDATE %s SET user_ip=$1, refresh_token=$2, time_life_rt=$3 WHERE id=$4 RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user.UserIP, user.RefreshToken, user.TimeLifeRT, user.ID)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
