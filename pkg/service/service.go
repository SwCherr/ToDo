package service

import "app/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
