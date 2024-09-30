package repository

import (
	"app"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// UPDATE таблица SET поле = значение
func (r *AuthPostgres) CreateSession(user_id int, user_ip, token string) error {
	var id int
	time := time.Now().Add(720 * time.Hour).Unix()

	query := fmt.Sprintf("UPDATE %s SET user_ip=$1, refresh_token=$2, time_life_rt=$3 WHERE id=$4 RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user_ip, token, time, user_id)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) GetUserById(id int) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *AuthPostgres) DeleteSession(user app.User) error {
	var id int
	query := fmt.Sprintf("UPDATE %s SET user_ip=$1, refresh_token=$2, time_life_rt=$3 WHERE id=$4 RETURNING id", userTable)
	user_row := r.db.QueryRow(query, "", "", "", user.ID)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}
