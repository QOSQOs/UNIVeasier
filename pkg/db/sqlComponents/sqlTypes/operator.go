package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type SQLOperator int8

const (
	UNKNOWN_OPERATOR SQLOperator = -1 + iota
	SELECT
	UPDATE
	DELETE
)

var SQLOperatorList = [...]string{
	"SELECT",
	"UPDATE",
	"DELETE",
}

func (query SQLOperator) IsValid() error {
	if query < 0 || int(query) >= len(SQLOperatorList) {
		return &errors.InvalidTypeIndexError{"SQL Operator", int8(query)}
	}
	return nil
}

func (query SQLOperator) ToString() (string, error) {
	if err := query.IsValid(); err == nil {
		return SQLOperatorList[query], nil
	} else {
		return "", err
	}
}

func ToSQLOperator(query string) (SQLOperator, error) {
	for i, val := range SQLOperatorList {
		if val == query {
			return SQLOperator(i), nil
		}
	}
	return UNKNOWN_OPERATOR, &errors.InvalidTypeError{"SQL Operator", query}
}
