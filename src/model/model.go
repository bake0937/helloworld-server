package model

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type QueryHandler interface {
	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type Model struct {
	db QueryHandler
}

func NewModel(db QueryHandler) *Model {
	return &Model{
		db: db,
	}
}

func (model *Model) Begin() (*Model, error) {
	db, ok := model.db.(*sqlx.DB)
	if !ok {
		return nil, errors.New("Failed to begin a transaction")
	}

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return NewModel(tx), nil
}

func (model *Model) Commit() error {
	tx, ok := model.db.(*sqlx.Tx)
	if !ok {
		return errors.New("Failed to commit this transaction")
	}
	return tx.Commit()
}

func (model *Model) Rollback() error {
	tx, ok := model.db.(*sqlx.Tx)
	if !ok {
		return errors.New("Failed to rollback this transaction")
	}
	return tx.Rollback()
}
