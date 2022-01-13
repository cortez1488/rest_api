package service

import (
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/cortez1488/rest_todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoServer.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos)}
}
