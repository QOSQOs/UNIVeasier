package model

import (
	"testing"

	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
	"github.com/stretchr/testify/assert"
)

func TestValidationPerson(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numberTest    int8
		firstName     string
		lastName      string
		email         string
		gender        types.Gender
		typePerson    types.Person
		isVerified    types.Status
		expectedError bool
	}{
		{1, "juan gabriel", "quispe soto", "juan@qosqo.com", types.MALE, types.STUDENT, types.UNVERIFIED, false},
		{2, "juana", "ramirez", "juana10@univeasier.net", types.FEMALE, types.VISITOR, types.PENDING, false},
		{3, "gabriel10", "quispe soto", "juan10@qosqo.com", types.MALE, types.GRADUATED, types.PENDING, true},
		{4, "pepe", "ramirez 123", "pepe@univeasier.net", types.MALE, types.GRADUATED, types.PENDING, true},
		{5, "pepe", "diaz", "a@pe-.12edu", types.MALE, types.GRADUATED, types.PENDING, true},
		{6, "luis", "quispe", "a@pe.edu", types.UNKNOWN_GENDER, types.VISITOR, types.PENDING, true},
		{7, "juan dany paul", "quispe lana", "a@pe.edu", types.FEMALE, types.UNKNOWN_PERSON, types.VERIFIED, true},
		{8, "gabriel soto", "castro", "gabi@pe.edu", types.FEMALE, types.UNKNOWN_PERSON, types.UNKNOWN_STATUS, true},
	}

	for i, test := range tests {
		personModel := Person{
			FirstName:  types.StringFrom(test.firstName),
			LastName:   types.StringFrom(test.lastName),
			Email:      types.StringFrom(test.email),
			Gender:     test.gender,
			Type:       test.typePerson,
			IsVerified: test.isVerified}

		err := personModel.Validate()
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}
