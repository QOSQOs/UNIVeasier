package sqlComponents

import (
	"fmt"

	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
)

type SQLFilter struct {
	logicalOperator sqlTypes.Logical
	compOperator    sqlTypes.Comparator
	sqlColumn       SQLColumn
}

func (filter *SQLFilter) IsValid() error {
	err := filter.logicalOperator.IsValid()
	if err != nil && filter.logicalOperator != sqlTypes.UNKNOWN_LOGICAL {
		return err
	}

	err = filter.compOperator.IsValid()
	if err != nil && filter.compOperator != sqlTypes.UNKNOWN_COMPARATOR {
		return err
	}

	return nil
}

func (filter *SQLFilter) ToString() (string, error) {
	if err := filter.IsValid(); err != nil {
		return "", err
	}

	withLogicalOperator := true
	logicalOp, err := filter.logicalOperator.ToString()
	if err != nil {
		withLogicalOperator = false
	}

	compOp, err := filter.compOperator.ToString()
	if err != nil {
		return "", err
	}

	columnName := filter.sqlColumn.GetName()
	columnValue := filter.sqlColumn.GetValue()

	sqlStringFilter := fmt.Sprintf("%s %s %s", columnName, compOp, columnValue)
	if withLogicalOperator {
		sqlStringFilter = fmt.Sprintf("%s %s", logicalOp, sqlStringFilter)
	}

	return sqlStringFilter, nil
}

func newFilter(logical sqlTypes.Logical, columnName string, cmp sqlTypes.Comparator, value interface{}) (SQLFilter, error) {
	sqlColumn := SQLColumn{name: columnName, value: value}

	sqlFilter := SQLFilter{
		logicalOperator: logical,
		compOperator:    cmp,
		sqlColumn:       sqlColumn,
	}

	if err := sqlFilter.IsValid(); err != nil {
		return sqlFilter, err
	}

	return sqlFilter, nil
}
