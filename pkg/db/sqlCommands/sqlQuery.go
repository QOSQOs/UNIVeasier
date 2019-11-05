package sqlCommands

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/common/marks"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlCommands/sqlTypes"

	"database/sql"
	"fmt"
	"time"
)

/*
addFilter(colum string, value interface{})
var SQLFilter(conn, table, queryType) validate querytype, table
Columns("a", "b", "c", "d", ...) validate columns

addOrFilter(a, >, 3) validate columns, op,
addAndFilter(b, =, 1,2,3,4)
addFilter
addOR
addSetValue(b, 2) validate column
addSetValue(c, false) validate column
*/

type SQLColumn struct {
	name   string
	value  interface{}
	enable bool
}

func (col *SQLColumn) IsUsed() bool {
	return col.enable
}

func (col *SQLColumn) ToString() string {
	switch value := col.value.(type) {
	case string:
		return fmt.Sprintf("%q", value)
	case bool:
		boolValue := fmt.Sprintf("%v", value)
		if boolValue == "true" {
			return "1"
		}
		return "0"
	case time.Time:
		return fmt.Sprintf("%q", value.Format(common.TimeFormat))
	default:
		return fmt.Sprintf("%v", value)
	}
}

type SQLFilter1 struct {
	logicalOperator sqlTypes.LogOperator
	compOperator    sqlTypes.SQLOperator
	sqlColumn       SQLColumn
}

func (filter *SQLFilter1) ToSTring() {
	return fmt.Sprintf("")
}

func newFilter(logicOp sqlTypes.LogOperator, colName string, compOp sqlTypes.SQLOperator, val interface{}) (SQLFilter1, error) {
	if err := logicOp.IsValid(); err != nil {
		return SQLFilter1{}, err
	}

	if err := compOp.IsValid(); err != nil {
		return SQLFilter1{}, err
	}

	sqlColumn := SQLColumn{name: colName, value: val}
	return SQLFilter1{
		logicalOperator: logicOp,
		compOperator:    compOp,
		sqlColumn:       sqlColumn,
	}, nil
}

type SQLQuery1 struct {
	tableName   string
	queryType   sqlTypes.SQLQueryTitle
	columnNames map[string]SQLColumn
	filters     []SQLFilter
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

func (query *SQLQuery1) AddColumns(columnNames ...string) error {
	for _, column := range columnNames {
		if sqlColumn, ok := query.columnNames[column]; !ok {
			sqlColumn.enable = true
			query.columnNames[column] = sqlColumn

		} else {
			return &errors.ValueNotExistError{column, "ColumnNames"}
		}
	}
	return nil
}

func (query *SQLQuery1) SetColumnValue(columnNames []string, values []interface{}) error {
	if len(columnNames) != len(values) {
		return &errors.NotEqualsSizeError{"column names", "values"}
	}

	for i, column := range columnNames {
		if sqlColumn, ok := query.columnNames[column]; !ok {
			sqlColumn.value = values[i]
			query.columnNames[column] = sqlColumn

		} else {
			return &errors.ValueNotExistError{column, "ColumnNames"}
		}
	}
	return nil
}

func NewQuery(conn *sql.DB, tableName, sqlQueryType string) (SQLQuery1, error) {
	queryType, err := sqlTypes.ToSQLQueryTitle(sqlQueryType)
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

type SQLQuery struct {
	DBName            string
	TableName         string
	QueryType         sqlTypes.SQLQueryTitle
	ColumnsName       map[string]bool
	FilterExpressions []SQLFilter
}

func (sql *SQLQuery) Init(conn *sql.DB, dbName string, tableName string) error {
	res, err := conn.Query("call dbqosqos.GetColumnByTableName(?)", tableName)
	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetColumnByTableName"), "info", err.Error())
		return err
	}

	sql.ColumnsName = make(map[string]bool)

	for res.Next() {
		var columnName string
		err = res.Scan(&columnName)
		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return err
		}
		sql.ColumnsName[columnName] = false
	}

	if len(sql.ColumnsName) == 0 {
		return &errors.TableNotExistError{tableName}
	}
	return nil
}

func (sql *SQLQuery) AddColumn(nameColumn string) error {
	if _, ok := sql.ColumnsName[nameColumn]; ok {
		sql.ColumnsName[nameColumn] = true
		return nil
	}
	return &errors.ValueNotExistError{nameColumn, "ColumnsName"}
}

func (sql *SQLQuery) AddFilter(sqlFilter SQLFilter) error {
	if err := sqlFilter.Op.IsValid(); err != nil {
		return err
	}

	if _, ok := sql.ColumnsName[sqlFilter.ColumnName]; !ok {
		return &errors.ValueNotExistError{sqlFilter.ColumnName, "ColumnsName"}
	}

	sql.FilterExpressions = append(sql.FilterExpressions, sqlFilter)
	return nil
}

func (sql *SQLQuery) GetFilterExpressions() (string, error) {
	finalFilterExpressions := marks.EMPTY
	for i, sqlFilter := range sql.FilterExpressions {
		filterExpression, err := sqlFilter.GetFilterExpression()
		if err != nil {
			return "", err
		}
		if i == 0 {
			finalFilterExpressions += filterExpression
		} else {
			linkString, _ := sqlFilter.Link.ToString()
			finalFilterExpressions += marks.SPACE + linkString + marks.SPACE + filterExpression
		}
	}
	return finalFilterExpressions, nil
}

func (sql *SQLQuery) GetSQLQuery() (string, error) {
	sqlTitle, err := sql.QueryType.ToString()
	if err != nil {
		return "", err
	}
	finalSQLQuery := sqlTitle + marks.SPACE

	switch sqlTitle := sql.QueryType; sqlTitle {
	case sqlTypes.SELECT:
		if sql.ColumnsName["ALL_COLUMNS"] {
			finalSQLQuery += marks.ASTERISK + marks.SPACE
		} else {
			firstColumn := true
			for columnName, needToBeIncluded := range sql.ColumnsName {
				if !needToBeIncluded {
					continue
				}
				if firstColumn {
					finalSQLQuery += columnName
					firstColumn = false
				} else {
					finalSQLQuery += marks.COMMA + marks.SPACE + columnName
				}
			}
			finalSQLQuery += marks.SPACE
		}
	case sqlTypes.DELETE:
		// nothing to do
	case sqlTypes.UPDATE:
		// TODO: implement the logic for th update query,
		// probably this requiere to do some changes in the struct
	default:
		sqlTitleString, _ := sqlTitle.ToString()
		return "", &errors.InvalidTypeError{"SQLTitle", sqlTitleString}
	}

	finalSQLQuery += "FROM" + marks.SPACE + sql.TableName + marks.SPACE + "WHERE" + marks.SPACE
	filterExpression, err := sql.GetFilterExpressions()
	if err != nil {
		return "", err
	}
	finalSQLQuery += filterExpression + marks.SEMICOLON
	return finalSQLQuery, nil
}
