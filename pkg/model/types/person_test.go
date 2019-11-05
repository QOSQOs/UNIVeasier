package types

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestValidatePersonType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	const PEOPLE Person = 100

	var tests = []struct {
		numberTest     int8
		person         Person
		expectedString string
		expectedError  bool
	}{
		{0, VISITOR, "VISITOR", false},
		{1, STUDENT, "STUDENT", false},
		{2, GRADUATED, "GRADUATED", false},
		{3, PROFESSOR, "PROFESSOR", false},
		{4, PEOPLE, "PEOPLE", true},
	}

	t.Run("person type to string", func(t *testing.T) {
		for i, test := range tests {
			personString, err := test.person.ToString()
			if !test.expectedError {
				assert.NoError(err, utils.FailedTest(i))
				assert.Equal(test.expectedString, personString, utils.FailedTest(i))
			}
		}
	})

	t.Run("validate person type", func(t *testing.T) {
		for i, test := range tests {
			err := test.person.IsValid()
			if test.expectedError {
				assert.Error(err, utils.FailedTest(i))
			} else {
				assert.NoError(err, utils.FailedTest(i))
			}
		}
	})
}

func TestConvertToPersonType(t *testing.T) {
	common.InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		numberTest     int8
		personString   string
		expectedPerson Person
		expectedError  bool
	}{
		{0, "VISITOR", VISITOR, false},
		{1, "STUDENT", STUDENT, false},
		{2, "GRADUATED", GRADUATED, false},
		{3, "PROFESSOR", PROFESSOR, false},
		{4, "HELLO", UNKNOWN_PERSON, true},
		{5, "NONE", UNKNOWN_PERSON, true},
	}

	for i, test := range tests {
		person, err := ToPerson(test.personString)
		if !test.expectedError {
			assert.NoError(err, utils.FailedTest(i))
		} else {
			assert.Error(err, utils.FailedTest(i))
		}
		assert.Equal(test.expectedPerson, person, utils.FailedTest(i))
	}
}
