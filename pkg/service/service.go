package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetPareToken(session app.Sesion) (acces, refresh string, err error)
	RefreshToken(session app.Sesion) (acces, refresh string, err error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
