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
	CreateList(userId int, list todoServer.TodoList) (int, error)
	GetList(listId int, userId int) (todoServer.TodoList, error)
	GetAllLists(userId int) ([]todoServer.TodoList, error)
	DeleteList(listId int, userId int) error
	UpdateList(listId, userId int, input todoServer.UpdateListInput) error
}
type TodoItem interface {
	CreateItem(userId, listId int, input todoServer.TodoItem) (int, error)
	GetItems(userId, listId int) ([]todoServer.TodoItem, error)
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos),
		TodoList: NewListService(repos.TodoList),
		TodoItem: NewItemService(repos.TodoItem, repos.TodoList),
	}
}
