package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(user, password, host, port, dbName string) (*sqlx.DB, error) {
	dsn := buildDSN(user, password, host, port, dbName)
	return sqlx.Connect("mysql", dsn)
}
