package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type SQLLogical int8

const (
	UNKNOWN_LOGICAL SQLLogical = -1 + iota
	AND
	OR
)

var SQLLogicalList = [...]string{
	"AND",
	"OR",
}

func (op SQLLogical) IsValid() error {
	if op < 0 || int(op) >= len(SQLLogicalList) {
		return &errors.InvalidTypeIndexError{"SQL Logical", int8(op)}
	}
	return nil
}

func (op SQLLogical) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return SQLLogicalList[op], nil
	} else {
		return "", err
	}
}

func ToSQLLogical(op string) (SQLLogical, error) {
	for i, val := range SQLLogicalList {
		if val == op {
			return SQLLogical(i), nil
		}
	}
	return UNKNOWN_LOGICAL, &errors.InvalidTypeError{"SQL Logical", op}
}
