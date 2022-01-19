package repository

import (
	"fmt"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type todoItemPostgres struct {
	db *sqlx.DB
}

func newTodoItemPostgres(db *sqlx.DB) *todoItemPostgres {
	return &todoItemPostgres{db: db}
}

func (r *todoItemPostgres) CreateItem(listId int, item todoServer.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	createTodoItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createTodoItemQuery, item.Title, item.Description)
	var itemId int
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createUsersListsQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *todoItemPostgres) GetItems(listId int) ([]todoServer.TodoItem, error) {
	var items []todoServer.TodoItem
	query := fmt.Sprintf("select i.id, i.title, i.description, i.done FROM %s i JOIN %s li on i.id = li.item_id where li.list_id = $1", todoItemsTable, listsItemsTable)
	//query := fmt.Sprintf("select i.id, i.title, i.description, i.done FROM %s i JOIN %s li on i.id = li.item_id" +
	//	" JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Select(&items, query, listId)
	if err != nil {
		return nil, err
	}
	return items, err
}

func (r *todoItemPostgres) GetById(userId, itemId int) (todoServer.TodoItem, error) {
	query := fmt.Sprintf("SELECT i.id, i.title, i.description, i.done FROM %s i JOIN %s li ON i.id = li.item_id JOIN %s ul ON ul.list_id = li.list_id WHERE i.id = $1 AND ul.user_id = $2 LIMIT 1",
		todoItemsTable, listsItemsTable, usersListsTable)
	var result todoServer.TodoItem
	if err := r.db.Get(&result, query, itemId, userId); err != nil {
		return todoServer.TodoItem{}, err
	}

	return result, nil
}

func (r *todoItemPostgres) UpdateItem(itemId, userId int, input todoServer.UpdateItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s i SET %s FROM %s li,  %s ul WHERE i.id = li.item_id AND ul.list_id = li.list_id AND i.id = $%d AND ul.user_id = $%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, itemId, userId)

	fmt.Printf("updateQuery: %s \n", query)
	fmt.Printf("args: %s \n", args)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
