package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	GetUserByGuid(id int) (app.User, error)
	GeneratePareTokens(user_id int, user_ip string) (acces, refresh string, err error)
	RefreshToken(user_id int, user_ip, token string) (acces, refresh string, err error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
