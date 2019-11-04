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
		{0, types.NullString(), types.Int64From(1), true},
		{1, types.StringFrom(""), types.Int64From(1), true},
		{2, types.StringFrom("IA"), types.NullInt64(), true},
		{3, types.StringFrom("D3.js v5"), types.Int64From(15), false},
		{4, types.StringFrom("IA"), types.Int64From(1), false},
		{5, types.StringFrom(strings.Repeat("D", 101)), types.Int64From(1), true},
		{6, types.StringFrom(strings.Repeat("D", 99)), types.Int64From(1), false},
	}

	for i, test := range tests {
		interestModel := Interest{
			Tag:       test.tag,
			CreatedBy: test.createdBy}

		err := interestModel.Validate()
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}
