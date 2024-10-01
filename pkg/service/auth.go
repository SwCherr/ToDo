package service

import (
	"app"
	"app/pkg/repository"
	"crypto/rand"
	"crypto/sha1"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt            = "ahdbvjdccjdn"
	siginkey        = "djdjdjbvhdn3d^&*(" // ключ подписи для расшифровки токена
	accessTokenITL  = 30 * time.Minute
	refreshTokenITL = 720 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	GUID   string `json:"guid"`
	UserIP string `json:"ip"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = s.generateHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetPareToken(input app.Sesion) (acces, refresh string, err error) {
	newAccessToken, err := s.generateAccessToken(input)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return "", "", err
	}

	input.RefreshToken = s.generateHash(newRefreshToken) // hash for save in db
	if err := s.createSession(input); err != nil {
		return "", "", err
	}

	// предача newRefreshToken в формате base64
	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
	return newAccessToken, newRefreshToken, nil
}

func (s *AuthService) RefreshToken(input app.Sesion) (access, refresh string, err error) {
	user_session, err := s.repo.PullOutSessionByGUID(input.GUID)
	if err != nil {
		return "", "", err
	}

	// декодирование RefreshToken из base64
	token, err := b64.StdEncoding.DecodeString(input.RefreshToken)
	if err != nil {
		return "", "", err
	}
	input.RefreshToken = string(token)

	if user_session.RefreshToken != s.generateHash(input.RefreshToken) {
		return "", "", errors.New("invalid refresh token")
	}

	if user_session.UserIP != input.UserIP {
		// send email
		user, err := s.repo.GetUserById(user_session.UserID)
		if err != nil {
			return "", "", err
		}
		s.sendEmail(user.Email)
		return "", "", errors.New("invalid IP addres")
	}

	if user_session.ExpiresIn < time.Now().Unix() {
		return "", "", errors.New("refresh token expired")
	}

	return s.GetPareToken(user_session)
}

func (s *AuthService) generateAccessToken(user_session app.Sesion) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenITL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user_session.GUID,
		user_session.UserIP})
	return token.SignedString([]byte(siginkey))
}

func (s *AuthService) generateRefreshToken() (string, error) {
	size_token := 64
	refresh_token := make([]byte, size_token)
	_, err := rand.Read(refresh_token)
	if err != nil {
		return "", err
	}
	return string(refresh_token), nil
}

func (s *AuthService) createSession(session app.Sesion) error {
	session.ExpiresIn = time.Now().Add(refreshTokenITL).Unix()
	return s.repo.CreateSession(session)
}

func (s *AuthService) generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
