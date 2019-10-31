package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidateGenderType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const LESBIAN Gender = 100

	var tests = []struct {
		numberTest     int8
		gender         Gender
		expectedString string
		expectedError  bool
	}{
		{0, MALE, "MALE", false},
		{1, FEMALE, "FEMALE", false},
		{2, OTHER, "OTHER", false},
		{3, LESBIAN, "LESBIAN", true},
	}

	t.Run("gender type to string", func(t *testing.T) {
		for i, test := range tests {
			genderString, err := test.gender.ToString()
			if !test.expectedError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, genderString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate gender type", func(t *testing.T) {
		for i, test := range tests {
			err := test.gender.IsValid()
			if test.expectedError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToGenderType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest     int8
		genderString   string
		expectedGender Gender
		expectedError  bool
	}{
		{0, "MALE", MALE, false},
		{1, "FEMALE", FEMALE, false},
		{2, "OTHER", OTHER, false},
		{3, "HELLO", UNKNOWN_GENDER, true},
		{4, "NONE", UNKNOWN_GENDER, true},
	}

	for i, test := range tests {
		gender, err := ToGender(test.genderString)
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedGender, gender, utils.FailedTest(i))
	}
}
