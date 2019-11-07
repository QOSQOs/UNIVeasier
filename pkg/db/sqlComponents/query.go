package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"

	"database/sql"
)

type SQLQuery1 struct {
	tableName   string
	queryType   sqlTypes.SQLOperator
	columnNames map[string]SQLColumn
	filters     []SQLFilter1
}

func (query *SQLQuery1) getColumns(conn *sql.DB, tableName string) error {
	res, err := conn.Query("call GetColumnByTableName(?)", tableName)
	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetColumnByTableName"), "info", err.Error())
		return err
	}

	query.columnNames = make(map[string]SQLColumn)

	for res.Next() {
		var columnName string
		err = res.Scan(&columnName)
		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return err
		}
		query.columnNames[columnName] = SQLColumn{name: columnName}
	}

	if len(query.columnNames) == 0 {
		return &errors.TableNotExistError{tableName}
	}

	return nil
}

func (query *SQLQuery1) AddColumnsHeader(columnNames ...string) error {
	for _, column := range columnNames {
		if sqlColumn, ok := query.columnNames[column]; ok {
			sqlColumn.enable = true
			query.columnNames[column] = sqlColumn
		} else {
			return &errors.ValueNotExistError{column, "ColumnNames"}
		}
	}
	return nil
}

func (query *SQLQuery1) SetColumnsValues(columnNames []string, values []interface{}) error {
	if len(columnNames) != len(values) {
		return &errors.NotEqualsSizeError{"column names", "values"}
	}

	for i, column := range columnNames {
		if sqlColumn, ok := query.columnNames[column]; ok {
			sqlColumn.value = values[i]
			sqlColumn.enable = true
			query.columnNames[column] = sqlColumn
		} else {
			return &errors.ValueNotExistError{column, "ColumnNames"}
		}
	}
	return nil
}

func (query *SQLQuery1) addFilter(logical sqlTypes.SQLLogical, columnName string, cmp sqlTypes.SQLComparator, value interface{}) error {
	if _, ok := query.columnNames[columnName]; !ok {
		return &errors.ValueNotExistError{columnName, "ColumnNames"}
	}

	sqlFilter, err := newFilter(logical, columnName, cmp, value)
	if err != nil {
		return err
	}

	query.filters = append(query.filters, sqlFilter)
	return nil
}

func (query *SQLQuery1) AddFilter(columnName string, cmp sqlTypes.SQLComparator, value interface{}) error {
	err := query.addFilter(sqlTypes.UNKNOWN_LOGICAL, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery1) AddANDFilter(columnName string, cmp sqlTypes.SQLComparator, value interface{}) error {
	err := query.addFilter(sqlTypes.AND, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery1) AddORFilter(columnName string, cmp sqlTypes.SQLComparator, value interface{}) error {
	err := query.addFilter(sqlTypes.OR, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery1) GetSQLQuery() (string, error) {
	var sqlQuery string

	switch query.queryType {
	case sqlTypes.SELECT:
		return sqlQuery, nil
	}

	return sqlQuery, nil
}

//TODO
// getheaders()
// getfilters()

func NewQuery(conn *sql.DB, tableName, sqlQueryType string) (SQLQuery1, error) {
	queryType, err := sqlTypes.ToSQLOperator(sqlQueryType)
	if err != nil {
		return SQLQuery1{}, err
	}

	query := SQLQuery1{
		tableName: tableName,
		queryType: queryType,
	}

	err = query.getColumns(conn, tableName)
	if err != nil {
		return SQLQuery1{}, err
	}

	return query, nil
}
