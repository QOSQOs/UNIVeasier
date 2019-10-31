package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateStatusType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const PROGRESS Status = 100

	var tests = []struct {
		numberTest     int8
		status         Status
		expectedString string
		expectedError  bool
	}{
		{0, UNVERIFIED, "UNVERIFIED", false},
		{1, PENDING, "PENDING", false},
		{2, VERIFIED, "VERIFIED", false},
		{3, PROGRESS, "PROGRESS", true},
	}

	t.Run("status type to string", func(t *testing.T) {
		for i, test := range tests {
			statusString, err := test.status.ToString()
			if !test.expectedError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, statusString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate status type", func(t *testing.T) {
		for i, test := range tests {
			err := test.status.IsValid()
			if test.expectedError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToStatusType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest     int8
		statusString   string
		expectedStatus Status
		expectedError  bool
	}{
		{1, "UNVERIFIED", UNVERIFIED, false},
		{2, "PENDING", PENDING, false},
		{3, "VERIFIED", VERIFIED, false},
		{4, "HELLO", UNKNOWN_STATUS, true},
		{5, "NONE", UNKNOWN_STATUS, true},
	}

	for i, test := range tests {
		status, err := ToStatus(test.statusString)
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedStatus, status, utils.FailedTest(i))
	}
}
