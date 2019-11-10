package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type Comparator int8

const (
	UNKNOWN_COMPARATOR Comparator = -1 + iota
	EQUAL
	NOT_EQUAL
	GREATER
	GREATER_EQUALS
	LESS
	LESS_EQUALS
	IN
	NOT_IN
)

var ComparatorList = [...]string{
	"=",
	"<>",
	">",
	">=",
	"<",
	"<=",
	"IN",
	"NOT IN",
}

func (op Comparator) IsValid() error {
	if op < 0 || int(op) >= len(ComparatorList) {
		return &errors.InvalidTypeIndexError{"SQL Comparator", int8(op)}
	}
	return nil
}

func (op Comparator) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return ComparatorList[op], nil
	} else {
		return "", err
	}
}

func ToComparator(op string) (Comparator, error) {
	for i, val := range ComparatorList {
		if val == op {
			return Comparator(i), nil
		}
	}
	return UNKNOWN_COMPARATOR, &errors.InvalidTypeError{"SQL Comparator", op}
}
