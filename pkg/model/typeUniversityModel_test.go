package model

import (
	"testing"

	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
	"github.com/stretchr/testify/assert"
)

func TestValidationTypeUniversity(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name          string
		description   string
		isVerified    common.InstanceStatus
		createdBy     types.Int64
		expectedError bool
	}{
		{"public", "public universities", 1, types.NewInt64(1, true), false},
		{"private", "private universities", 2, types.NewInt64(2, true), false},
		{"private", "private universities", 2, types.NewInt64(0, false), true},
		{"state", "it is an invalid example", 1, types.NewInt64(0, false), true},
		{"", "", 10, types.NewInt64(1, true), true},
	}

	for _, t := range tests {
		model := TypeUniversity{
			Name:        types.StringFrom(t.name),
			Description: types.StringFrom(t.description),
			IsVerified:  t.isVerified,
			CreatedBy:   t.createdBy}

		err := model.Validate()
		if t.expectedError {
			assert.Error(err, model)
		} else {
			assert.NoError(err, model)
		}
	}
}
