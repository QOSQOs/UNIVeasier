package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceStatus(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		instanceStatus  InstanceStatus
		expectedResult1 string
		expectedResult2 bool
	}{
		{1, "UNVERIFIED", true},
		{2, "PENDING", true},
		{3, "VERIFIED", true},
		{4, "", false},
		{5, "", false},
	}
	for _, test := range tests {
		name := test.instanceStatus.String()
		assert.Equal(name, test.expectedResult1)
		err := test.instanceStatus.IsValid()
		assert.Equal((err == nil), test.expectedResult2)
	}
}

func TestInstanceStatusGetByName(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		name           string
		expectedResult InstanceStatus
	}{
		{"UNVERIFIED", 1},
		{"PENDING", 2},
		{"VERIFIED", 3},
		{"HELLO", 0},
		{"NONE", 0},
	}
	for _, test := range tests {
		var model InstanceStatus
		instanceStatus, _ := model.GetByName(test.name)
		assert.Equal(instanceStatus, test.expectedResult)
	}
}
