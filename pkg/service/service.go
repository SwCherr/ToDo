package service

import (
	todo "app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)                                 // return id new user in db & error
	CreateSession(username, password, token, user_ip string) error          // return error
	GenerateAccessToken(username, password, user_ip string) (string, error) // return generated token & error
	GenerateRefreshToken() (string, error)                                  // return generated token & error
	// ParseAccsessToken(token string) (int, error)                            // return user`s id from db & error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
