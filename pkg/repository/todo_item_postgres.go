package repository

import (
	"fmt"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
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
