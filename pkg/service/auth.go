package service

import (
	todo "app"
	"app/pkg/repository"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt            = "ahdbvjdccjdn"      // good practice
	siginkey        = "djdjdjbvhdn3d^&*(" // ключ подписи, нужен для расшифровки токена
	accessTokenITL  = 1 * time.Hour
	refreshTokenITL = 720 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	UserIP string `json:"user_ip"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordaHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) CreateSession(username, password, token, user_ip string) error {
	user, err := s.repo.GetUser(username, s.generatePasswordaHash(password))
	if err != nil {
		return err
	}
	user.UserIP = user_ip
	user.RefreshToken = token
	user.TimeLifeRT = time.Now().Add(accessTokenITL).Unix()
	return s.repo.CreateSession(user)
}

func (s *AuthService) generatePasswordaHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateAccessToken(username, password, user_ip string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, s.generatePasswordaHash(password))
	if err != nil {
		return "", err
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenITL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user_ip})
	return token.SignedString([]byte(siginkey))
}

func (s *AuthService) GenerateRefreshToken() (string, error) {
	size_token := 64
	refresh_token := make([]byte, size_token)
	_, err := rand.Read(refresh_token)
	if err != nil {
		return "", err
	}
	return s.generatePasswordaHash(string(refresh_token)), nil
}

// func (s *AuthService) ParseAccsessToken(accessToken string) (int, error) {
// 	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid signing method")
// 		}
// 		return []byte(siginkey), nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	claims, ok := token.Claims.(*tokenClaims)
// 	if !ok {
// 		return 0, errors.New("token claims are not of type *tokenClaims")
// 	}

// 	return claims.UserId, nil
// }
