package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/internal/common/marks"
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
)

type SQLFilter struct {
	Link            sqlTypes.SQLLogical
	ColumnName      string
	Op              sqlTypes.SQLComparator
	Values          []string
	ValuesAreString bool
}

func (sql *SQLFilter) GetFilterExpression() (string, error) {
	var sqlExpressionString = marks.EMPTY
	opString, _ := sql.Op.ToString()

	switch op := sql.Op; op {
	case
		sqlTypes.EQUAL,
		sqlTypes.NOT_EQUAL,
		sqlTypes.GREATER,
		sqlTypes.GREATER_EQUALS,
		sqlTypes.LESS,
		sqlTypes.LESS_EQUALS:

		sqlExpressionString = sql.ColumnName + marks.SPACE + opString + marks.SPACE + sql.GetWordFrom(sql.Values[0])
	case
		sqlTypes.IN,
		sqlTypes.NOT_IN:

		sqlExpressionString = sql.ColumnName + marks.SPACE + opString + marks.SPACE + marks.OPEN_BRACKET
		for i, value := range sql.Values {
			if i == 0 {
				sqlExpressionString += sql.GetWordFrom(value)
			} else {
				sqlExpressionString += marks.COMMA + marks.SPACE + sql.GetWordFrom(value)
			}
		}
		sqlExpressionString += marks.CLOSE_BRACKET
	default:
		opString, _ = op.ToString()
		return "", &errors.InvalidTypeError{"operator", opString}
	}
	return sqlExpressionString, nil
}

func (sql *SQLFilter) GetWordFrom(word string) string {
	if sql.ValuesAreString {
		return marks.QUOTE + word + marks.QUOTE
	} else {
		return word
	}
}
