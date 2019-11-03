package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type SQLOperator int8

const (
	UNKNOWN_OPERATOR SQLOperator = -1 + iota
	EQUAL
	NOT_EQUAL
	GREATER
	GREATER_EQUALS
	LESS
	LESS_EQUALS
	IN
	NOT_IN
)

var SQLOperatorList = [...]string{
	"=",
	"<>",
	">",
	">=",
	"<",
	"<=",
	"IN",
	"NOT IN",
}

func (op SQLOperator) IsValid() error {
	if op < 0 || int(op) >= len(SQLOperatorList) {
		return &errors.InvalidTypeIndexError{"operator", int8(op)}
	}
	return nil
}

func (op SQLOperator) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return SQLOperatorList[op], nil
	} else {
		return "", err
	}
}

func ToOperator(op string) (SQLOperator, error) {
	for i, val := range SQLOperatorList {
		if val == op {
			return SQLOperator(i), nil
		}
	}
	return UNKNOWN_OPERATOR, &errors.InvalidTypeError{"operator", op}
}
