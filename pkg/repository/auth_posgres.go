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

func (r *AuthPostgres) CreateSession(user_id int, user_ip, token string) error {
	var id int
	time := time.Now().Add(720 * time.Hour).Unix()

	query := fmt.Sprintf("INSERT INTO %s (guid, ip, token, time) values ($1, $2, $3, $4) RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user_id, user_ip, token, time)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUserByGuid(id int) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE guid=$1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *AuthPostgres) DeleteSession(user app.User) error {
	var id int
	query := fmt.Sprintf("UPDATE %s SET ip=$1, token=$2, time=$3 WHERE guid=$4 RETURNING id", userTable)
	user_row := r.db.QueryRow(query, "", "", "", user.ID)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}
