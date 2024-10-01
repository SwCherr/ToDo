package repository

import (
	"app"
	"fmt"

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
	query := fmt.Sprintf("INSERT INTO %s (email, password) values ($1, $2) RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user.Email, user.Password)
	if err := user_row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserById(id int) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *AuthPostgres) PullOutSessionByGUID(GUID string) (app.Sesion, error) {
	var session app.Sesion
	query := fmt.Sprintf("SELECT * FROM %s WHERE guid=$1", sessionTable)
	err := r.db.Get(&session, query, GUID)

	var id int
	delete := fmt.Sprintf("DELETE FROM %s WHERE guid=$1 RETURNING id", sessionTable)
	user_row := r.db.QueryRow(delete, GUID)
	if err := user_row.Scan(&id); err != nil {
		return session, err
	}
	return session, err

}

func (r *AuthPostgres) CreateSession(s app.Sesion) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, ip, guid, refresh_token, expires_in) values ($1, $2, $3, $4, $5) RETURNING id", sessionTable)
	user_row := r.db.QueryRow(query, s.UserID, s.UserIP, s.GUID, s.RefreshToken, s.ExpiresIn)
	if err := user_row.Scan(&id); err != nil {
		return err
	}
	return nil
}
