package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMarshalInt64Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest      int8
		number          Int64
		expectedMarshal string
	}{
		{0, Int64From(10), `10`},
		{1, NullInt64(), `null`},
	}

	for i, test := range tests {
		marshal, err := test.number.MarshalJSON()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedMarshal, string(marshal), utils.FailedTest(i))
	}
}

func TestUnmarshalInt64Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest        int8
		buffer            string
		expectedUnmarshal Int64
		expectedError     bool
	}{
		{0, `10`, Int64From(10), false},
		{1, `9223372036854775807`, Int64From(9223372036854775807), false},
		{2, `9223372036854775808`, NullInt64(), true},
		{3, `null`, NullInt64(), false},
		{4, `"null"`, NullInt64(), true},
		{5, `"10"`, NullInt64(), true},
	}

	for i, test := range tests {
		var unmarshal Int64
		err := unmarshal.UnmarshalJSON([]byte(test.buffer))
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
			assert.Equal(test.expectedUnmarshal, unmarshal, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
	}
}
