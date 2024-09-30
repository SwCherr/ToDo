package service

import (
	"app"
	"app/pkg/repository"
	"crypto/rand"
	"crypto/sha1"
	"errors"
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

func (s *AuthService) GetUserByGuid(id int) (app.User, error) {
	return s.repo.GetUserByGuid(id)
}

func (s *AuthService) generatePasswordaHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateAccessToken(user_id int, user_ip string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenITL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user_id,
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

func (s *AuthService) GeneratePareTokens(user_id int, user_ip string) (acces, refresh string, err error) {
	acces_token, err := s.GenerateAccessToken(user_id, user_ip)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := s.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	if err := s.createSession(user_id, user_ip, refresh_token); err != nil {
		return "", "", err
	}
	return acces_token, refresh_token, nil
}

func (s *AuthService) createSession(user_id int, user_ip, token string) error {
	return s.repo.CreateSession(user_id, user_ip, token)
}

func (s *AuthService) RefreshToken(user_id int, user_ip, token string) (access, refresh string, err error) {
	user, err := s.repo.GetUserByGuid(user_id)
	if err != nil {
		return "", "", err
	}

	if user.UserIP != user_ip {
		s.repo.DeleteSession(user)
		// send email ----------------------------------------------
		return "", "", errors.New("invalid IP addres: 1 - " + user_ip + " 2 - " + user.UserIP)
	}

	if user.RefreshToken != token {
		s.repo.DeleteSession(user)
		return "", "", errors.New("invalid refresh token")
	}

	if user.TimeLifeRT < time.Now().Unix() {
		s.repo.DeleteSession(user)
		return "", "", errors.New("refresh token expired")
	}
	return s.GeneratePareTokens(user_id, user_ip)
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
