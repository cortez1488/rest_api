package repository

import (
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoServer.User) (int, error)
	GetUser(username, password string) (todoServer.User, error)
}

type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthDb(db)}
}
