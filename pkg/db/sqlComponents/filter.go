package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"

	"fmt"
)

type SQLFilter1 struct {
	logicalOperator sqlTypes.SQLLogical
	compOperator    sqlTypes.SQLComparator
	sqlColumn       SQLColumn
}

func (filter *SQLFilter1) IsValid() error {
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

func (filter *SQLFilter1) ToSTring() (string, error) {
	if err := filter.IsValid(); err != nil {
		return "", err
	}

	withLogicalOperator := false
	logicalOp, err := filter.logicalOperator.ToString()
	if err != nil {
		withLogicalOperator = true
	}

	compOp, err := filter.compOperator.ToString()
	if err != nil {
		return "", err
	}

	columnName := filter.sqlColumn.ToString()
	columnValue := filter.sqlColumn.GetValue()

	sqlFilter := fmt.Sprintf("%s %s %s", columnName, compOp, columnValue)
	if withLogicalOperator {
		sqlFilter = fmt.Sprintf("%s %s", logicalOp, sqlFilter)
	}

	return sqlFilter, nil
}

func newFilter(logical sqlTypes.SQLLogical, columnName string, cmp sqlTypes.SQLComparator, value interface{}) (SQLFilter1, error) {
	sqlColumn := SQLColumn{name: columnName, value: value}

	sqlFilter := SQLFilter1{
		logicalOperator: logical,
		compOperator:    cmp,
		sqlColumn:       sqlColumn,
	}

	if err := sqlFilter.IsValid(); err != nil {
		return sqlFilter, err
	}

	return sqlFilter, nil
}
