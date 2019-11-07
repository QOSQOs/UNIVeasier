package sqlComponents

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetFilterExpression(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numberTest            int8
		column                string
		valuesAreString       bool
		op                    sqlTypes.SQLComparator
		values                []string
		expectedSQLExpression string
	}{
		{0, "column1", false, sqlTypes.EQUAL, []string{"1"}, "column1 = 1"},
		{1, "column2", true, sqlTypes.NOT_EQUAL, []string{"2"}, "column2 <> '2'"},
		{2, "column3", true, sqlTypes.GREATER, []string{"3"}, "column3 > '3'"},
		{3, "column4", false, sqlTypes.GREATER_EQUALS, []string{"4"}, "column4 >= 4"},
		{4, "column5", false, sqlTypes.IN, []string{"1", "2", "3"}, "column5 IN (1, 2, 3)"},
		{5, "column6", true, sqlTypes.IN, []string{"1", "2", "3"}, "column6 IN ('1', '2', '3')"},
		{6, "column7", true, sqlTypes.NOT_IN, []string{"1", "2"}, "column7 NOT IN ('1', '2')"},
	}

	for i, test := range tests {
		sqlFilter := SQLFilter{
			ColumnName:      test.column,
			ValuesAreString: test.valuesAreString,
			Op:              test.op,
			Values:          test.values}

		filterExpresion, _ := sqlFilter.GetFilterExpression()
		assert.Equal(test.expectedSQLExpression, filterExpresion, utils.FailedTest(i))
	}
}
