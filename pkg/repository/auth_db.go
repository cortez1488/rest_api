package repository

import (
	"fmt"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/jmoiron/sqlx"
)

type AuthDb struct {
	db *sqlx.DB
}

func NewAuthDb(db *sqlx.DB) *AuthDb {
	return &AuthDb{db: db}
}

func (r *AuthDb) CreateUser(user todoServer.User) (int, error) { // dividing DB logic
	// do some DB logic

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthDb) GetUser(username, password string) (todoServer.User, error) {
	var user todoServer.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
