package repository

import (
	"fmt"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListLite struct {
	db *sqlx.DB
}

func NewTodoListLite(db *sqlx.DB) *TodoListLite {
	return &TodoListLite{db: db}
}

func (r *TodoListLite) CreateList(userId int, list todoServer.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	createTodoListsQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createTodoListsQuery, list.Title, list.Description)
	var id int
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListLite) GetList(listId int, userId int) (todoServer.TodoList, error) {
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE tl.id = $1 AND ul.user_id = $2", todoListsTable, usersListsTable)
	var list todoServer.TodoList
	err := r.db.Get(&list, query, listId, userId)
	return list, err
}

func (r *TodoListLite) GetAllLists(userId int) ([]todoServer.TodoList, error) {
	var lists []todoServer.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListLite) DeleteList(listId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.list_id = $1 AND ul.user_id = $2", todoListsTable, usersListsTable) //don't work
	_, err := r.db.Exec(query, listId, userId)
	return err
}

func (r *TodoListLite) UpdateList(listId, userId int, input todoServer.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	if err != nil {

	}
	return err
}
