package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateLogical(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const NOT Logical = 100

	var tests = []struct {
		numberTest     int8
		logical        Logical
		expectedString string
		wantError      bool
	}{
		{0, AND, "AND", false},
		{1, OR, "OR", false},
		{2, NOT, "NOT", true},
	}

	t.Run("logical type to string", func(t *testing.T) {
		for i, test := range tests {
			logicalString, err := test.logical.ToString()
			if !test.wantError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, logicalString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate logical type", func(t *testing.T) {
		for i, test := range tests {
			err := test.logical.IsValid()
			if test.wantError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToLogicalType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest      int8
		logicalString   string
		expectedLogical Logical
		wantError       bool
	}{

		{0, "AND", AND, false},
		{1, "OR", OR, false},
		{7, "NOT", UNKNOWN_LOGICAL, true},
	}

	for i, test := range tests {
		logical, err := ToLogical(test.logicalString)
		if test.wantError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedLogical, logical, utils.FailedTest(i))
	}
}
