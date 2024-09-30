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

func (r *AuthPostgres) GetUserByGuid(guid int) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE guid=$1", userTable)
	err := r.db.Get(&user, query, guid)
	return user, err
}

func (r *AuthPostgres) CreateSession(user_guid int, user_ip, token string) error {
	var id int
	time := time.Now().Add(720 * time.Hour).Unix() // перенести задание времени в сервисы ????
	email := "optika.space@gmail.com"
	query := fmt.Sprintf("INSERT INTO %s (guid, ip, token, time, email) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user_guid, user_ip, token, time, email)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) UpdateSession(user_guid int, user_ip, token string) error {
	var id int
	time := time.Now().Add(720 * time.Hour).Unix() // перенести задание времени в сервисы ????
	query := fmt.Sprintf("UPDATE %s SET token=$1, time=$2 WHERE guid=$3 RETURNING id", userTable)
	user_row := r.db.QueryRow(query, token, time, user_guid)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}
