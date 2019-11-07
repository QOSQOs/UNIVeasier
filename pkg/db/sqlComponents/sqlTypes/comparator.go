package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type SQLComparator int8

const (
	UNKNOWN_COMPARATOR SQLComparator = -1 + iota
	EQUAL
	NOT_EQUAL
	GREATER
	GREATER_EQUALS
	LESS
	LESS_EQUALS
	IN
	NOT_IN
)

var SQLComparatorList = [...]string{
	"=",
	"<>",
	">",
	">=",
	"<",
	"<=",
	"IN",
	"NOT IN",
}

func (op SQLComparator) IsValid() error {
	if op < 0 || int(op) >= len(SQLComparatorList) {
		return &errors.InvalidTypeIndexError{"SQL Comparator", int8(op)}
	}
	return nil
}

func (op SQLComparator) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return SQLComparatorList[op], nil
	} else {
		return "", err
	}
}

func ToComparator(op string) (SQLComparator, error) {
	for i, val := range SQLComparatorList {
		if val == op {
			return SQLComparator(i), nil
		}
	}
	return UNKNOWN_COMPARATOR, &errors.InvalidTypeError{"SQL Comparator", op}
}
