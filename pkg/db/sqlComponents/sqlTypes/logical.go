package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type Logical int8

const (
	UNKNOWN_LOGICAL Logical = -1 + iota
	AND
	OR
)

var LogicalList = [...]string{
	"AND",
	"OR",
}

func (op Logical) IsValid() error {
	if op < 0 || int(op) >= len(LogicalList) {
		return &errors.InvalidTypeIndexError{"SQL Logical", int8(op)}
	}
	return nil
}

func (op Logical) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return LogicalList[op], nil
	} else {
		return "", err
	}
}

func ToLogical(op string) (Logical, error) {
	for i, val := range LogicalList {
		if val == op {
			return Logical(i), nil
		}
	}
	return UNKNOWN_LOGICAL, &errors.InvalidTypeError{"SQL Logical", op}
}
