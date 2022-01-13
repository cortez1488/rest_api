package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

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
