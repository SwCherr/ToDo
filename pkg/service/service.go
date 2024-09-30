package service

import (
	"app/pkg/repository"
)

type Authorization interface {
	GeneratePareTokens(user_id int, user_ip string) (acces, refresh string, err error)
	RefreshToken(user_id int, user_ip, token string) (acces, refresh string, err error)
	CreateSession(user_guid int, user_ip, token string) error
	UpdateSession(user_id int, user_ip, token string) error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
