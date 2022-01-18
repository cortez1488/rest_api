package repository

import (
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user todoServer.User) (int, error)
	GetUser(username, password string) (todoServer.User, error)
}

type TodoList interface {
	CreateList(userId int, list todoServer.TodoList) (int, error)
	GetList(listId int, userId int) (todoServer.TodoList, error)
	GetAllLists(userId int) ([]todoServer.TodoList, error)
	DeleteList(listId int, userId int) error
	UpdateList(listId, userId int, input todoServer.UpdateListInput) error
}
type TodoItem interface {
	CreateItem(listId int, input todoServer.TodoItem) (int, error)
	GetItems(listId int) ([]todoServer.TodoItem, error)
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthDb(db),
		TodoList: NewTodoListLite(db),
		TodoItem: newTodoItemPostgres(db)}
}
