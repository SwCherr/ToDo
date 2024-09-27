package service

import (
	todo "app"
	"app/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "ahdbvjdccjdn" // good practice

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordaHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordaHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
