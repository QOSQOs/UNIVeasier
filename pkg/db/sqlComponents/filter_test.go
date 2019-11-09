package sqlComponents

import (
	"testing"

	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents/sqlTypes"
	"github.com/stretchr/testify/assert"
)

func TestSQLFilterIsValid(t *testing.T) {
	assert := assert.New(t)

	var testsFilters = []struct {
		numberTest int8
		logical    sqlTypes.Logical
		comparator sqlTypes.Comparator
		wantError  bool
	}{
		{0, sqlTypes.OR, sqlTypes.GREATER, false},
		{1, sqlTypes.UNKNOWN_LOGICAL, sqlTypes.LESS, false},
		{2, sqlTypes.AND, sqlTypes.UNKNOWN_COMPARATOR, false},
		// These cases are added since it is allowed
		// to send any number as a type of comparator
		{3, 10, sqlTypes.LESS, true},
		{4, sqlTypes.OR, 20, true},
	}

	for i, test := range testsFilters {
		_, err := newFilter(test.logical, "", test.comparator, "")
		if test.wantError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}

func TestSQLFilterToString(t *testing.T) {
	assert := assert.New(t)

	var testsFilters = []struct {
		numberTest     int8
		wantError      bool
		equals         bool
		expectedString string
		filter         SQLFilter
	}{
		{
			numberTest:     0,
			wantError:      false,
			equals:         true,
			expectedString: `A > 10`,
			filter: SQLFilter{
				compOperator: sqlTypes.GREATER,
				sqlColumn:    SQLColumn{name: "A", value: 10},
			},
		},
		{
			numberTest:     1,
			wantError:      false,
			equals:         true,
			expectedString: `A = "test"`,
			filter: SQLFilter{
				compOperator: sqlTypes.EQUAL,
				sqlColumn:    SQLColumn{name: "A", value: "test"},
			},
		},
		{
			numberTest:     2,
			wantError:      false,
			equals:         true,
			expectedString: `A = 1`,
			filter: SQLFilter{
				compOperator: sqlTypes.EQUAL,
				sqlColumn:    SQLColumn{name: "A", value: true},
			},
		},
		{
			numberTest:     3,
			wantError:      false,
			equals:         false,
			expectedString: `A = "2"`,
			filter: SQLFilter{
				compOperator: sqlTypes.EQUAL,
				sqlColumn:    SQLColumn{name: "A", value: 2},
			},
		},
		{
			numberTest:     3,
			wantError:      false,
			equals:         false,
			expectedString: `B = "2"`,
			filter: SQLFilter{
				compOperator: sqlTypes.EQUAL,
				sqlColumn:    SQLColumn{name: "A", value: "2"},
			},
		},
		{
			numberTest:     3,
			wantError:      true,
			equals:         false,
			expectedString: `A = 1`,
			filter: SQLFilter{
				compOperator: sqlTypes.UNKNOWN_COMPARATOR,
				sqlColumn:    SQLColumn{name: "A", value: true},
			},
		},
	}

	for i, test := range testsFilters {
		sqlStringFilter, err := test.filter.ToString()
		if test.wantError {
			assert.Error(err, utils.FailedTest(i), utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i), utils.FailedTest(i))
			if test.equals {
				assert.Equal(test.expectedString, sqlStringFilter, utils.FailedTest(i))
			} else {
				assert.NotEqual(test.expectedString, sqlStringFilter, utils.FailedTest(i))
			}
		}
	}
}
