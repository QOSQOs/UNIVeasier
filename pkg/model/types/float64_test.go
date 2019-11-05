package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMarshalFloat64Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest      int8
		number          Float64
		expectedMarshal string
	}{
		{0, Float64From(10.2), `10.2`},
		{1, NullFloat64(), `null`},
	}

	for i, test := range tests {
		marshal, err := test.number.MarshalJSON()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedMarshal, string(marshal), utils.FailedTest(i))
	}
}

func TestUnmarshalFloat64Type(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest        int8
		buffer            string
		expectedUnmarshal Float64
		expectedError     bool
	}{
		{0, `10.2`, Float64From(10.2), false},
		{1, `1.797693134862315708145274237317043567981e+308`, Float64From(1.797693134862315708145274237317043567981e+308), false},
		{2, `1.797693134862315708145274237317043567981e+309`, NullFloat64(), true},
		{3, `null`, NullFloat64(), false},
		{4, `"null"`, NullFloat64(), true},
		{5, `"10.2"`, NullFloat64(), true},
	}

	for i, test := range tests {
		var unmarshal Float64
		err := unmarshal.UnmarshalJSON([]byte(test.buffer))
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
			assert.Equal(test.expectedUnmarshal, unmarshal, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
	}
}
