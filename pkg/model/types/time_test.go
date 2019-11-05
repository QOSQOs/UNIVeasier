package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestMarshalTimeType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	timeTest0, err := time.Parse(timeFormat, "2019-11-01T15:04:05Z")
	assert.NoError(err)

	var tests = []struct {
		numberTest      int8
		timestamp       Time
		expectedMarshal string
	}{
		{0, TimeFrom(timeTest0), `"2019-11-01T15:04:05Z"`},
		{1, NullTime(), `null`},
	}

	for i, test := range tests {
		marshal, err := test.timestamp.MarshalJSON()
		assert.NoError(err, utils.FailedTest(i))
		assert.Equal(test.expectedMarshal, string(marshal), utils.FailedTest(i))
	}
}

func TestUnmarshalTimeType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	timeTest0, err := time.Parse(timeFormat, "2019-11-01T15:04:05Z")
	assert.NoError(err)

	var tests = []struct {
		numberTest        int8
		buffer            string
		expectedUnmarshal Time
		expectedError     bool
	}{
		{0, `"2019-11-01T15:04:05Z"`, TimeFrom(timeTest0), false},
		{1, `null`, NullTime(), false},
		{2, `"2019-19-19T15:04:05Z"`, NullTime(), true},
		{3, `"my date"`, NullTime(), true},
	}

	for i, test := range tests {
		var unmarshal Time
		err := unmarshal.UnmarshalJSON([]byte(test.buffer))
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
			assert.Equal(test.expectedUnmarshal, unmarshal, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
	}
}
