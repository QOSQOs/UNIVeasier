package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/config"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/connection"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	configPath := "C:/Users/I504018/go/src/github.com/QOSQOs/UNIVeasier/internal/config/config.json"
	configVars, _ := config.ReadConfiguration(configPath)
	conn, _ := connection.InitConnection(configVars)

	sqlQueryModel := SQLQuery{}

	var testsTable = []struct {
		numberTest    int8
		dbName        string
		tableName     string
		expectedError bool
	}{
		{0, "dbqosqos", "person", false},
		{1, "dbqosqos", "dbqosqos", true},
	}

	for i, test := range testsTable {
		err := sqlQueryModel.Init(conn, test.dbName, test.tableName)
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}

func TestAddColumn(t *testing.T) {
	assert := assert.New(t)
	sqlQueryModel := SQLQuery{
		DBName:      "qosqos",
		TableName:   "person",
		ColumnsName: map[string]bool{"id": false, "code": false}}

	var testsColumns = []struct {
		numberTest    int8
		columnName    string
		expectedError bool
	}{
		{0, "id", false},
		{1, "code", false},
		{2, "column", true},
	}

	for i, test := range testsColumns {
		err := sqlQueryModel.AddColumn(test.columnName)
		fmt.Println(err)
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}

func TestAddFilter(t *testing.T) {
	assert := assert.New(t)
	sqlQueryModel := SQLQuery{
		DBName:      "qosqos",
		TableName:   "person",
		ColumnsName: map[string]bool{"id": false, "code": false}}

	testsFilters := []struct {
		numberTest    int8
		filter        SQLFilter
		expectedError bool
	}{
		{0, SQLFilter{
			ColumnName:      "id",
			Op:              sqlTypes.EQUAL,
			Values:          []string{"1"},
			ValuesAreString: true}, false},
		{1, SQLFilter{
			ColumnName:      "column",
			Op:              sqlTypes.NOT_IN,
			Values:          []string{"1", "2"},
			ValuesAreString: false}, true},
		{2, SQLFilter{
			ColumnName:      "column",
			Op:              sqlTypes.UNKNOWN_COMPARATOR,
			Values:          []string{"1", "2"},
			ValuesAreString: false}, true},
	}

	for i, test := range testsFilters {
		err := sqlQueryModel.AddFilter(test.filter)
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}

func TestGetFilterExpressions(t *testing.T) {
	assert := assert.New(t)
	sqlQueryModel := SQLQuery{
		DBName:      "qosqos",
		TableName:   "person",
		ColumnsName: map[string]bool{"id": false, "code": false}}

	testsFilters := []struct {
		numberTest               int8
		filters                  []SQLFilter
		expectedFilterExpression string
	}{
		{0, []SQLFilter{
			SQLFilter{
				Link:            sqlTypes.UNKNOWN_LOGICAL,
				ColumnName:      "id",
				Op:              sqlTypes.EQUAL,
				Values:          []string{"1"},
				ValuesAreString: true},
			SQLFilter{
				Link:            sqlTypes.AND,
				ColumnName:      "code",
				Op:              sqlTypes.NOT_IN,
				Values:          []string{"1", "2"},
				ValuesAreString: false}},
			"id = '1' AND code NOT IN (1, 2)"},
	}

	for i, test := range testsFilters {
		for _, filter := range test.filters {
			_ = sqlQueryModel.AddFilter(filter)
		}

		filterExpression, _ := sqlQueryModel.GetFilterExpressions()
		assert.Equal(test.expectedFilterExpression, filterExpression, utils.FailedTest(i))
	}
}

func TestGetSQLQuery(t *testing.T) {
	assert := assert.New(t)
	sqlQueryModel := SQLQuery{
		DBName:    "qosqos",
		TableName: "person"}

	sqlFilters := []SQLFilter{
		SQLFilter{
			Link:            sqlTypes.UNKNOWN_LOGICAL,
			ColumnName:      "id",
			Op:              sqlTypes.EQUAL,
			Values:          []string{"1"},
			ValuesAreString: true},
		SQLFilter{
			Link:            sqlTypes.AND,
			ColumnName:      "code",
			Op:              sqlTypes.NOT_IN,
			Values:          []string{"1", "2"},
			ValuesAreString: false}}

	tests := []struct {
		numberTest       int8
		queryType        sqlTypes.SQLOperator
		columnsName      map[string]bool
		expectedSQLQuery string
	}{
		{0, sqlTypes.SELECT, map[string]bool{"id": true, "code": false, "first_name": true},
			"SELECT id, first_name FROM person WHERE id = '1' AND code NOT IN (1, 2);"},
		{1, sqlTypes.SELECT, map[string]bool{"id": true, "code": false, "ALL_COLUMNS": true},
			"SELECT * FROM person WHERE id = '1' AND code NOT IN (1, 2);"},
		{2, sqlTypes.DELETE, map[string]bool{"id": true, "code": false, "first_name": true},
			"DELETE FROM person WHERE id = '1' AND code NOT IN (1, 2);"},
		{3, sqlTypes.DELETE, map[string]bool{"id": true, "code": true, "first_name": false},
			"DELETE FROM person WHERE id = '1' AND code NOT IN (1, 2);"},
	}

	for i, test := range tests {
		sqlQueryModel.FilterExpressions = sqlFilters
		sqlQueryModel.QueryType = test.queryType
		sqlQueryModel.ColumnsName = test.columnsName

		sqlQuery, _ := sqlQueryModel.GetSQLQuery()
		assert.Equal(test.expectedSQLQuery, sqlQuery, utils.FailedTest(i))
	}
}

func TestSQLColumnToString(t *testing.T) {
	assert := assert.New(t)

	timeTest, err := time.Parse(common.TimeFormat, "2019-11-01T15:04:05Z")
	assert.NoError(err)

	var testsColumns = []struct {
		numberTest    int8
		value         interface{}
		expectedValue string
	}{
		{0, 1, `1`},
		{1, 1.5, `1.5`},
		{2, "value", `"value"`},
		{3, true, `1`},
		{4, false, `0`},
		{5, timeTest, `"2019-11-01T15:04:05Z"`},
	}

	for i, test := range testsColumns {
		sqlColumn := SQLColumn{value: test.value}
		stringValue := sqlColumn.ToString()
		assert.Equal(test.expectedValue, stringValue, utils.FailedTest(i))
	}
}
