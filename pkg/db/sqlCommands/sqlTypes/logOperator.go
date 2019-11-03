package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type LogOperator int8

const (
	UNKNOWN_LOG_OPERATOR LogOperator = -1 + iota
	AND
	OR
)

var LogOperatorList = [...]string{
	"AND",
	"OR",
}

func (op LogOperator) IsValid() error {
	if op < 0 || int(op) >= len(LogOperatorList) {
		return &errors.InvalidTypeIndexError{"operator", int8(op)}
	}
	return nil
}

func (op LogOperator) ToString() (string, error) {
	if err := op.IsValid(); err == nil {
		return LogOperatorList[op], nil
	} else {
		return "", err
	}
}

func ToLogOperator(op string) (LogOperator, error) {
	for i, val := range LogOperatorList {
		if val == op {
			return LogOperator(i), nil
		}
	}
	return UNKNOWN_LOG_OPERATOR, &errors.InvalidTypeError{"logOperator", op}
}
