package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUserById(id int) (app.User, error)
	GeneratePareTokens(user_id int, user_ip string) (acces, refresh string, err error)
	RefreshToken(user_id int, user_ip, token string) (acces, refresh string, err error)
	// ParseAccsessToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
