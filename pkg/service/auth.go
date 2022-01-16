package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/cortez1488/rest_todo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "ihateniggers1488"
	signingKey = "65465213498sajkgfkjsafjlk###fsaf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todoServer.User) (int, error) {
	// do some SERIVICE job
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user) // and call it do DB
}

func (s *AuthService) generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}, user.Id})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of *tokenClaims")
	}
	return claims.UserId, nil
}
