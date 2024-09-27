package service

import (
	todo "app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error) // return id new user in db & error
}

type TodoList interface {
}

type TodoItems interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
