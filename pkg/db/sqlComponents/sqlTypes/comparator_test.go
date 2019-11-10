package sqlTypes

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateComparator(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const EXIST Comparator = 100

	var tests = []struct {
		numberTest     int8
		comparator     Comparator
		expectedString string
		wantError      bool
	}{
		{0, EQUAL, "=", false},
		{1, NOT_EQUAL, "<>", false},
		{2, GREATER, ">", false},
		{3, GREATER_EQUALS, ">=", false},
		{4, LESS, "<", false},
		{5, LESS_EQUALS, "<=", false},
		{6, IN, "IN", false},
		{7, NOT_IN, "NOT IN", false},
		{7, EXIST, "EXIST", true},
	}

	t.Run("comparator type to string", func(t *testing.T) {
		for i, test := range tests {
			comparatorString, err := test.comparator.ToString()
			if !test.wantError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, comparatorString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate comparator type", func(t *testing.T) {
		for i, test := range tests {
			err := test.comparator.IsValid()
			if test.wantError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToComparatorType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest         int8
		comparatorString   string
		expectedComparator Comparator
		wantError          bool
	}{

		{0, "=", EQUAL, false},
		{1, "<>", NOT_EQUAL, false},
		{2, ">", GREATER, false},
		{3, ">=", GREATER_EQUALS, false},
		{4, "<", LESS, false},
		{5, "<=", LESS_EQUALS, false},
		{6, "IN", IN, false},
		{7, "NOT IN", NOT_IN, false},
		{7, "EXIST", UNKNOWN_COMPARATOR, true},
	}

	for i, test := range tests {
		comparator, err := ToComparator(test.comparatorString)
		if test.wantError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedComparator, comparator, utils.FailedTest(i))
	}
}
