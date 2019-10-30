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
		{0, "UNVERIFIED", true},
		{1, "PENDING", true},
		{2, "VERIFIED", true},
		{3, "", false},
		{4, "", false},
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
		{"UNVERIFIED", 0},
		{"PENDING", 1},
		{"VERIFIED", 2},
		{"HELLO", -1},
		{"NONE", -1},
	}
	for _, test := range tests {
		var model InstanceStatus
		instanceStatus, _ := model.GetByName(test.name)
		assert.Equal(instanceStatus, test.expectedResult)
	}
}
