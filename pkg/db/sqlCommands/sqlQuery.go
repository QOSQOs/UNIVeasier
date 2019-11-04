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
addfilter(a, >, 3) validate columns, op,
addfilter(b, =, 1)
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

type SQLQuery1 struct {
	tableName   string
	queryType   sqlTypes.SQLQueryTitle
	columnNames map[string]SQLColumn
	filters     []SQLFilter
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
