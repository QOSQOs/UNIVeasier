package model

import (
	"testing"

	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		firstName     string
		lastName      string
		email         string
		gender        common.Gender
		typePerson    common.TypePerson
		isVerified    common.InstanceStatus
		expectedError bool
	}{
		{"juan gabriel", "quispe soto", "juan10@qosqo.com", 0, 0, 0, false},
		{"pepe", "ramirez", "pepe@univeasier.net", 1, 1, 1, false},
		{"juan gabriel", "quispe soto", "juan10@qosqo.com", 6, 7, 0, true},
		{"pepe", "ramirez", "pepe@univeasier.net", 1, 2, -2, true},
		{"", "diaz", "a@pe.edu", -1, 2, 5, true},
		{"luis10", "quispe", "a@pe.edu", 6, 7, 8, true},
		{"juan dany paul", "quispe lana", "a@pe.edu", 2, 2, 2, false},
		{"gabriel soto", "", "gabi@pe.edu1", -1, -2, -3, true},
	}

	for _, t := range tests {
		model := Person{
			FirstName:  types.StringFrom(t.firstName),
			LastName:   types.StringFrom(t.lastName),
			Email:      types.StringFrom(t.email),
			Gender:     t.gender,
			Type:       t.typePerson,
			IsVerified: t.isVerified}

		err := model.Validate()
		if t.expectedError {
			assert.Error(err, model)
		} else {
			assert.NoError(err, model)
		}
	}
}
