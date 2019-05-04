package model

import (
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
