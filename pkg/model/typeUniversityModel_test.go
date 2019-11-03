package model

import (
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationTypeUniversity(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		testNumber    int8
		name          string
		description   string
		isVerified    types.Status
		createdBy     types.Int64
		expectedError bool
	}{
		{0, "public", "public universities", types.UNVERIFIED, types.Int64From(1), false},
		{1, "private", "private universities", types.PENDING, types.Int64From(2), false},
		{2, "state", "it is an invalid example", types.UNKNOWN_STATUS, types.Int64From(1), true},
		{3, "private", "private universities", types.VERIFIED, types.NullInt64(), true},
	}

	for i, test := range tests {
		typeUniversityModel := TypeUniversity{
			Name:        types.StringFrom(test.name),
			Description: types.StringFrom(test.description),
			IsVerified:  test.isVerified,
			CreatedBy:   test.createdBy}

		err := typeUniversityModel.Validate()
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}
