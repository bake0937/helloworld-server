package database

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(user, password, host, port, dbName string) (*sqlx.DB, error) {
	dsn := buildDSN(user, password, host, port, dbName)
	return sqlx.Connect("mysql", dsn)
}

func buildDSN(user, password, host, port, dbName string) string {
	config := mysql.NewConfig()
	config.User = user
	config.Passwd = password
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%s", host, port)
	config.DBName = dbName
	config.ParseTime = true
	config.Params = map[string]string{
		"charset": "utf8mb4",
	}
	return config.FormatDSN()
}
