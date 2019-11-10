package sqlComponents

import (
	"testing"
	"time"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSQLColumnGetValue(t *testing.T) {
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
		stringValue := sqlColumn.GetValue()
		assert.Equal(test.expectedValue, stringValue, utils.FailedTest(i))
	}
}
