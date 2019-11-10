package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type Operator int8

const (
	UNKNOWN_OPERATOR Operator = -1 + iota
	SELECT
	UPDATE
	DELETE
)

var OperatorList = [...]string{
	"SELECT",
	"UPDATE",
	"DELETE",
}

func (query Operator) IsValid() error {
	if query < 0 || int(query) >= len(OperatorList) {
		return &errors.InvalidTypeIndexError{"SQL Operator", int8(query)}
	}
	return nil
}

func (query Operator) ToString() (string, error) {
	if err := query.IsValid(); err == nil {
		return OperatorList[query], nil
	} else {
		return "", err
	}
}

func ToOperator(query string) (Operator, error) {
	for i, val := range OperatorList {
		if val == query {
			return Operator(i), nil
		}
	}
	return UNKNOWN_OPERATOR, &errors.InvalidTypeError{"SQL Operator", query}
}
