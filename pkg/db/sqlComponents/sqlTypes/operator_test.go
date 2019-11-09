package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateOperator(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const REMOVE Operator = 100

	var tests = []struct {
		numberTest     int8
		operator       Operator
		expectedString string
		wantError      bool
	}{
		{0, SELECT, "SELECT", false},
		{1, UPDATE, "UPDATE", false},
		{2, DELETE, "DELETE", false},
		{3, REMOVE, "REMOVE", true},
	}

	t.Run("operator type to string", func(t *testing.T) {
		for i, test := range tests {
			operatorString, err := test.operator.ToString()
			if !test.wantError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, operatorString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate operator type", func(t *testing.T) {
		for i, test := range tests {
			err := test.operator.IsValid()
			if test.wantError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToOperatorType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest       int8
		operatorString   string
		expectedOperator Operator
		wantError        bool
	}{

		{0, "SELECT", SELECT, false},
		{1, "UPDATE", UPDATE, false},
		{2, "DELETE", DELETE, false},
		{3, "REMOVE", UNKNOWN_OPERATOR, true},
	}

	for i, test := range tests {
		operator, err := ToOperator(test.operatorString)
		if test.wantError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedOperator, operator, utils.FailedTest(i))
	}
}
