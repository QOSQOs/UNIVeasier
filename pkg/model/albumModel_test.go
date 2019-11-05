package model

import (
	"strings"
	"testing"

	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
	"github.com/stretchr/testify/assert"
)

func TestValidationAlbum(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		numberTest    int8
		Name          types.String
		Description   types.String
		CreatedBy     types.Int64
		expectedError bool
	}{
		{1, types.StringFrom("Album Informática 2018-I"), types.StringFrom("album de informática del año 2018"), types.Int64From(1), false},
		{2, types.StringFrom("Album Medicina 2015-II"), types.StringFrom("album de medicina del año 2015"), types.NullInt64(), true},
		{3, types.NullString(), types.StringFrom("album description test"), types.Int64From(1), true},
		{4, types.StringFrom(""), types.StringFrom("album description test"), types.Int64From(1), true},
		{5, types.StringFrom("Album Informática 2018-I"), types.NullString(), types.Int64From(1), true},
		{6, types.StringFrom("Album Informática 2018-I"), types.StringFrom(""), types.Int64From(1), true},
		{7, types.StringFrom(strings.Repeat("a", 99)), types.StringFrom("album description test"), types.Int64From(1), false},
		{8, types.StringFrom(strings.Repeat("a", 101)), types.StringFrom("album description test"), types.Int64From(1), true},
	}

	for i, test := range tests {
		albumModel := Album{
			Name:        test.Name,
			Description: test.Description,
			CreatedBy:   test.CreatedBy}

		err := albumModel.Validate()
		if test.expectedError {
			assert.Error(err, utils.FailedTest(i))
		} else {
			assert.NoError(err, utils.FailedTest(i))
		}
	}
}
