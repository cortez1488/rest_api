package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func NewSqliteDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "E:\\SQL\\SQLiteStudio\\rest_test.db")
	if err != nil {
		log.Fatal("error on opening the database: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error on ping: ", err.Error())
	}
	return db, err
}
