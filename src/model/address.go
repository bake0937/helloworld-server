package model

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (model *Model) AllFindAddress() ([]Address, error) {
	sql, args, err := sq.Select("id, email").From("address").ToSql()

	res := []Address{}
	err = model.db.Select(&res, sql, args...)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (model *Model) CreateAddress(address *Address) error {
	fmt.Printf("CreateAddress")

	model.Begin()

	sql, args, err := sq.Insert("address").Columns("email").Values(address.Email).ToSql()
	_, err2 := model.db.Exec(sql, args...)

	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}

	model.Commit()

	return nil
}
