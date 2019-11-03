package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
)

type SQLQueryTitle int8

const (
	UNKNOWN_SQLQUERY SQLQueryTitle = -1 + iota
	SELECT
	UPDATE
	DELETE
)

var SQLQueryTitleList = [...]string{
	"SELECT",
	"UPDATE",
	"DELETE",
}

func (query SQLQueryTitle) IsValid() error {
	if query < 0 || int(query) >= len(SQLQueryTitleList) {
		return &errors.InvalidTypeIndexError{"SQLQuery", int8(query)}
	}
	return nil
}

func (query SQLQueryTitle) ToString() (string, error) {
	if err := query.IsValid(); err == nil {
		return SQLQueryTitleList[query], nil
	} else {
		return "", err
	}
}

func ToSQLQueryTitle(query string) (SQLQueryTitle, error) {
	for i, val := range SQLQueryTitleList {
		if val == query {
			return SQLQueryTitle(i), nil
		}
	}
	return UNKNOWN_SQLQUERY, &errors.InvalidTypeError{"operator", query}
}
