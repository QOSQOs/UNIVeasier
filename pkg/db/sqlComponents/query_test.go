package sqlComponents

import (
	"testing"

	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
	"github.com/stretchr/testify/assert"
)

func getColumnsFromDataBaseMock(columnsTable []string) map[string]SQLColumn {
	mapColumns := make(map[string]SQLColumn)
	for _, column := range columnsTable {
		sqlColumn := SQLColumn{name: column, enable: false}
		mapColumns[column] = sqlColumn
	}
	return mapColumns
}

func TestSQLQueryValidGetSelectQuery(t *testing.T) {
	assert := assert.New(t)

	tableName := "TEST"
	tableColumns := []string{"A", "B", "C", "D", "E", "F"}

	var testsFilters = []struct {
		numberTest     int8
		columnHeaders  []string
		columnFilters  []string
		typeFilters    []sqlTypes.Logical
		comparators    []sqlTypes.Comparator
		values         []interface{}
		expectedString string
	}{
		{
			numberTest:     0,
			columnHeaders:  []string{"A", "B"},
			columnFilters:  []string{"C", "D", "E"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL, sqlTypes.AND, sqlTypes.OR},
			comparators:    []sqlTypes.Comparator{sqlTypes.GREATER, sqlTypes.EQUAL, sqlTypes.LESS_EQUALS},
			values:         []interface{}{2, true, "value"},
			expectedString: `SELECT A, B FROM TEST WHERE C > 2 AND D = 1 OR E <= "value"`,
		},
		{
			numberTest:     1,
			columnHeaders:  []string{"A", "B"},
			columnFilters:  []string{"A"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{2},
			expectedString: `SELECT A, B FROM TEST WHERE A = 2`,
		},
		{
			numberTest:     2,
			columnFilters:  []string{"B"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{"value"},
			expectedString: `SELECT * FROM TEST WHERE B = "value"`,
		},
		{
			numberTest:     3,
			columnHeaders:  []string{"A", "B", "C"},
			expectedString: `SELECT A, B, C FROM TEST`,
		},
	}

	for i, test := range testsFilters {
		mapColumns := getColumnsFromDataBaseMock(tableColumns)

		sqlQuery := SQLQuery{
			tableName:   tableName,
			queryType:   sqlTypes.SELECT,
			columnNames: mapColumns,
		}

		for _, header := range test.columnHeaders {
			err := sqlQuery.AddColumnsHeader(header)
			assert.NoError(err, utils.FailedTest(i))
		}

		for i, _ := range test.columnFilters {
			var err error
			switch test.typeFilters[i] {
			case sqlTypes.UNKNOWN_LOGICAL:
				err = sqlQuery.AddFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.AND:
				err = sqlQuery.AddANDFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.OR:
				err = sqlQuery.AddORFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			}
			assert.NoError(err, utils.FailedTest(i))
		}

		stringSqlQuery, err := sqlQuery.GetSQLQuery()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedString, stringSqlQuery, utils.FailedTest(i))
	}
}

func TestSQLQueryValidGetDeleteQuery(t *testing.T) {
	assert := assert.New(t)

	tableName := "TEST"
	tableColumns := []string{"A", "B", "C", "D", "E", "F"}

	var testsFilters = []struct {
		numberTest     int8
		columnFilters  []string
		typeFilters    []sqlTypes.Logical
		comparators    []sqlTypes.Comparator
		values         []interface{}
		expectedString string
	}{
		{
			numberTest:     0,
			columnFilters:  []string{"C", "D", "E"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL, sqlTypes.AND, sqlTypes.OR},
			comparators:    []sqlTypes.Comparator{sqlTypes.GREATER, sqlTypes.EQUAL, sqlTypes.LESS_EQUALS},
			values:         []interface{}{2, true, "value"},
			expectedString: `DELETE FROM TEST WHERE C > 2 AND D = 1 OR E <= "value"`,
		},
		{
			numberTest:     1,
			columnFilters:  []string{"A"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{2},
			expectedString: `DELETE FROM TEST WHERE A = 2`,
		},
		{
			numberTest:     2,
			columnFilters:  []string{"B"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{"value"},
			expectedString: `DELETE FROM TEST WHERE B = "value"`,
		},
		{
			numberTest:     3,
			expectedString: `DELETE FROM TEST`,
		},
	}

	for i, test := range testsFilters {
		mapColumns := getColumnsFromDataBaseMock(tableColumns)

		sqlQuery := SQLQuery{
			tableName:   tableName,
			queryType:   sqlTypes.DELETE,
			columnNames: mapColumns,
		}

		for i, _ := range test.columnFilters {
			var err error
			switch test.typeFilters[i] {
			case sqlTypes.UNKNOWN_LOGICAL:
				err = sqlQuery.AddFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.AND:
				err = sqlQuery.AddANDFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.OR:
				err = sqlQuery.AddORFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			}
			assert.NoError(err, utils.FailedTest(i))
		}

		stringSqlQuery, err := sqlQuery.GetSQLQuery()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedString, stringSqlQuery, utils.FailedTest(i))
	}
}

func TestSQLQueryValidGetUpdateQuery(t *testing.T) {
	assert := assert.New(t)

	tableName := "TEST"
	tableColumns := []string{"A", "B", "C", "D", "E", "F"}

	var testsFilters = []struct {
		numberTest     int8
		columnHeaders  []string
		valuesHeaders  []interface{}
		columnFilters  []string
		typeFilters    []sqlTypes.Logical
		comparators    []sqlTypes.Comparator
		values         []interface{}
		expectedString string
	}{
		{
			numberTest:     0,
			columnHeaders:  []string{"A", "B"},
			valuesHeaders:  []interface{}{2, 4},
			columnFilters:  []string{"C", "D", "E"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL, sqlTypes.AND, sqlTypes.OR},
			comparators:    []sqlTypes.Comparator{sqlTypes.GREATER, sqlTypes.EQUAL, sqlTypes.LESS_EQUALS},
			values:         []interface{}{2, true, "value"},
			expectedString: `UPDATE TEST SET A = 2, B = 4 WHERE C > 2 AND D = 1 OR E <= "value"`,
		},
		{
			numberTest:     1,
			columnHeaders:  []string{"A", "B"},
			valuesHeaders:  []interface{}{2, 4},
			columnFilters:  []string{"C"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{2},
			expectedString: `UPDATE TEST SET A = 2, B = 4 WHERE C = 2`,
		},
		{
			numberTest:     2,
			columnHeaders:  []string{"A"},
			valuesHeaders:  []interface{}{"value"},
			columnFilters:  []string{"B"},
			typeFilters:    []sqlTypes.Logical{sqlTypes.UNKNOWN_LOGICAL},
			comparators:    []sqlTypes.Comparator{sqlTypes.EQUAL},
			values:         []interface{}{"2"},
			expectedString: `UPDATE TEST SET A = "value" WHERE B = "2"`,
		},
		{
			numberTest:     3,
			columnHeaders:  []string{"B", "C", "D"},
			valuesHeaders:  []interface{}{"test", 1, "2"},
			expectedString: `UPDATE TEST SET B = "test", C = 1, D = "2"`,
		},
	}

	for i, test := range testsFilters {
		mapColumns := getColumnsFromDataBaseMock(tableColumns)

		sqlQuery := SQLQuery{
			tableName:   tableName,
			queryType:   sqlTypes.UPDATE,
			columnNames: mapColumns,
		}

		err := sqlQuery.SetColumnsValues(test.columnHeaders, test.valuesHeaders)
		assert.NoError(err, utils.FailedTest(i))

		for i, _ := range test.columnFilters {
			var err error
			switch test.typeFilters[i] {
			case sqlTypes.UNKNOWN_LOGICAL:
				err = sqlQuery.AddFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.AND:
				err = sqlQuery.AddANDFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			case sqlTypes.OR:
				err = sqlQuery.AddORFilter(test.columnFilters[i], test.comparators[i], test.values[i])
				break
			}
			assert.NoError(err, utils.FailedTest(i))
		}

		stringSqlQuery, err := sqlQuery.GetSQLQuery()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedString, stringSqlQuery, utils.FailedTest(i))
	}
}
