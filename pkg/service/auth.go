package service

import (
	todo "app"
	"app/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "ahdbvjdccjdn"      // good practice
	siginkey = "djdjdjbvhdn3d^&*(" // ключ подписи, нужен для расшифровки токена
	tokenITL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, s.generatePasswordaHash(password))

	if err != nil {
		return "", err
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{ // change on 512 !!!!
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenITL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID})

	return token.SignedString([]byte(siginkey))
}
