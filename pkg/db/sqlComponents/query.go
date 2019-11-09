package sqlComponents

import (
	"fmt"
	"strings"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"

	"database/sql"
)

type SQLQuery struct {
	tableName   string
	queryType   sqlTypes.Operator
	columnNames map[string]SQLColumn
	filters     []SQLFilter
}

func (query *SQLQuery) getColumns(conn *sql.DB, tableName string) error {
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

func (query *SQLQuery) AddColumnsHeader(columnNames ...string) error {
	if query.queryType != sqlTypes.SELECT {
		queryString, err := query.queryType.ToString()
		if err != nil {
			return err
		}
		return &errors.OperationNotSupportedError{"AddColumnHeader", queryString}
	}

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

func (query *SQLQuery) SetColumnsValues(columnNames []string, values []interface{}) error {
	if query.queryType != sqlTypes.UPDATE {
		queryString, err := query.queryType.ToString()
		if err != nil {
			return err
		}
		return &errors.OperationNotSupportedError{"SetColumnsValues", queryString}
	}

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

func (query *SQLQuery) addFilter(logical sqlTypes.Logical, columnName string, cmp sqlTypes.Comparator, value interface{}) error {
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

func (query *SQLQuery) AddFilter(columnName string, cmp sqlTypes.Comparator, value interface{}) error {
	if len(query.filters) != 0 {
		return &errors.InvalidFirstFilterError{}
	}
	err := query.addFilter(sqlTypes.UNKNOWN_LOGICAL, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery) AddANDFilter(columnName string, cmp sqlTypes.Comparator, value interface{}) error {
	if len(query.filters) == 0 {
		return &errors.InvalidFilterError{}
	}
	err := query.addFilter(sqlTypes.AND, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery) AddORFilter(columnName string, cmp sqlTypes.Comparator, value interface{}) error {
	if len(query.filters) == 0 {
		return &errors.InvalidFilterError{}
	}
	err := query.addFilter(sqlTypes.OR, columnName, cmp, value)
	if err != nil {
		return err
	}
	return nil
}

func (query *SQLQuery) getHeaders() (string, error) {
	headerExpression := ""

	switch query.queryType {
	case sqlTypes.SELECT:
		for _, column := range query.columnNames {
			if column.IsUsed() {
				headerExpression = fmt.Sprintf("%s %s,", headerExpression, column.GetName())
			}
		}
		if len(headerExpression) == 0 {
			return "*", nil
		}
		headerExpression = headerExpression[:len(headerExpression)-1]
	case sqlTypes.UPDATE:
		for _, column := range query.columnNames {
			if column.IsUsed() {
				headerExpression = fmt.Sprintf("%s %s = %s,", headerExpression, column.GetName(), column.GetValue())
			}
		}
		if len(headerExpression) == 0 {
			return "", &errors.EmptyHeaderExpressionError{}
		}
		headerExpression = headerExpression[:len(headerExpression)-1]
	case sqlTypes.DELETE:
		for _, column := range query.columnNames {
			if column.IsUsed() {
				return "", &errors.InvalidHeaderExpressionError{}
			}
		}
	}

	return strings.TrimLeft(headerExpression, " "), nil
}

func (query *SQLQuery) getFilters() (string, error) {
	filterExpression := ""

	for _, filter := range query.filters {
		stringFilter, err := filter.ToString()
		if err != nil {
			return "", err
		}
		filterExpression = fmt.Sprintf("%s %s", filterExpression, stringFilter)
	}

	if len(filterExpression) != 0 {
		filterExpression = fmt.Sprintf("WHERE %s", strings.TrimLeft(filterExpression, " "))
	}

	return filterExpression, nil
}

func (query *SQLQuery) getSelectQuery() (string, error) {
	selectString, err := query.queryType.ToString()
	if err != nil {
		return "", err
	}

	stringHeaders, err := query.getHeaders()
	if err != nil {
		return "", err
	}
	selectString = fmt.Sprintf("%s %s", selectString, stringHeaders)

	selectString = fmt.Sprintf("%s FROM %s", selectString, query.tableName)

	stringFilters, err := query.getFilters()
	if err != nil {
		return "", err
	}
	selectString = fmt.Sprintf("%s %s", selectString, stringFilters)

	return strings.TrimRight(selectString, " "), nil
}

func (query *SQLQuery) getUpdateQuery() (string, error) {
	updateString, err := query.queryType.ToString()
	if err != nil {
		return "", err
	}

	updateString = fmt.Sprintf("%s %s SET", updateString, query.tableName)

	stringHeaders, err := query.getHeaders()
	if err != nil {
		return "", err
	}
	updateString = fmt.Sprintf("%s %s", updateString, stringHeaders)

	stringFilters, err := query.getFilters()
	if err != nil {
		return "", err
	}
	updateString = fmt.Sprintf("%s %s", updateString, stringFilters)

	return strings.TrimRight(updateString, " "), nil
}

func (query *SQLQuery) getDeleteQuery() (string, error) {
	deleteString, err := query.queryType.ToString()
	if err != nil {
		return "", err
	}

	deleteString = fmt.Sprintf("%s FROM %s", deleteString, query.tableName)

	_, err = query.getHeaders()
	if err != nil {
		return "", err
	}

	stringFilters, err := query.getFilters()
	if err != nil {
		return "", err
	}
	deleteString = fmt.Sprintf("%s %s", deleteString, stringFilters)

	return strings.TrimRight(deleteString, " "), nil
}

func (query *SQLQuery) GetSQLQuery() (string, error) {
	var err error
	var sqlQueryString string

	switch query.queryType {
	case sqlTypes.SELECT:
		sqlQueryString, err = query.getSelectQuery()
	case sqlTypes.UPDATE:
		sqlQueryString, err = query.getUpdateQuery()
	case sqlTypes.DELETE:
		sqlQueryString, err = query.getDeleteQuery()
	}

	if err != nil {
		return "", err
	}
	return sqlQueryString, nil
}

func NewQuery(conn *sql.DB, tableName, sqlQueryType string) (SQLQuery, error) {
	queryType, err := sqlTypes.ToOperator(sqlQueryType)
	if err != nil {
		return SQLQuery{}, err
	}

	query := SQLQuery{
		tableName: tableName,
		queryType: queryType,
	}

	err = query.getColumns(conn, tableName)
	if err != nil {
		return SQLQuery{}, err
	}

	return query, nil
}
