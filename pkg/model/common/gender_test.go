package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		gender          Gender
		expectedResult1 string
		expectedResult2 bool
	}{
		{1, "MALE", true},
		{2, "FEMALE", true},
		{3, "OTHER", true},
		{4, "", false},
	}
	for _, test := range tests {
		name := test.gender.String()
		assert.Equal(name, test.expectedResult1)
		err := test.gender.IsValid()
		assert.Equal((err == nil), test.expectedResult2)
	}
}

func TestGenderGetByName(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		name           string
		expectedResult Gender
	}{
		{"MALE", 1},
		{"FEMALE", 2},
		{"OTHER", 3},
		{"HELLO", 0},
		{"NONE", 0},
	}
	for _, test := range tests {
		var model Gender
		gender, _ := model.GetByName(test.name)
		assert.Equal(gender, test.expectedResult)
	}
}
