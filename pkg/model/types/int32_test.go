package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMarshalInt32Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest      int8
		number          Int32
		expectedMarshal string
	}{
		{0, Int32From(10), `10`},
		{1, NullInt32(), `null`},
	}

	for i, test := range tests {
		marshal, err := test.number.MarshalJSON()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedMarshal, string(marshal), utils.FailedTest(i))
	}
}

func TestUnmarshalInt32Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest        int8
		buffer            string
		expectedUnmarshal Int32
		expectedError     bool
	}{
		{0, `10`, Int32From(10), false},
		{1, `2147483647`, Int32From(2147483647), false},
		{2, `2147483648`, NullInt32(), true},
		{3, `null`, NullInt32(), false},
		{4, `"null"`, NullInt32(), true},
		{5, `"10"`, NullInt32(), true},
	}

	for i, test := range tests {
		var unmarshal Int32
		err := unmarshal.UnmarshalJSON([]byte(test.buffer))
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
			assert.Equal(test.expectedUnmarshal, unmarshal, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
	}
}
