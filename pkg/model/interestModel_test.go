package model

import (
	"strings"
	"testing"

	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
	"github.com/stretchr/testify/assert"
)

func TestValidationInterest(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numberTest    int8
		tag           types.String
		createdBy     types.Int64
		expectedError bool
	}{
		{1, types.StringFrom("intership"), types.Int64From(1), false},
		{2, types.StringFrom("intership"), types.NullInt64(), true},
		{3, types.NullString(), types.Int64From(1), true},
		{4, types.StringFrom(""), types.Int64From(1), true},
		{5, types.StringFrom(strings.Repeat("a", 99)), types.Int64From(1), false},
		{6, types.StringFrom(strings.Repeat("a", 101)), types.Int64From(1), true},
	}

	for i, test := range tests {
		interesModel := Interest{
			Tag:       test.tag,
			CreatedBy: test.createdBy}

		err := interesModel.Validate()
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}
