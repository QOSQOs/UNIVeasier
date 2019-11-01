package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMarshalStringType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest      int8
		text            String
		expectedMarshal string
	}{
		{0, StringFrom("text"), `"text"`},
		{1, NullString(), `null`},
	}

	for i, test := range tests {
		marshal, err := test.text.MarshalJSON()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedMarshal, string(marshal), utils.FailedTest(i))
	}
}

func TestUnmarshalStringType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest        int8
		buffer            string
		expectedUnmarshal String
		expectedError     bool
	}{
		{0, `"bar"`, StringFrom("bar"), false},
		{1, `"foo10"`, StringFrom("foo10"), false},
		{2, `null`, NullString(), false},
		{3, `10`, NullString(), true},
		{4, `true`, NullString(), true},
	}

	for i, test := range tests {
		var unmarshal String
		err := unmarshal.UnmarshalJSON([]byte(test.buffer))
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
			assert.Equal(test.expectedUnmarshal, unmarshal, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
	}
}
