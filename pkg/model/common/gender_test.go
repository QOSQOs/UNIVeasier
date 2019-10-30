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
		{0, "MALE", true},
		{1, "FEMALE", true},
		{2, "OTHER", true},
		{3, "", false},
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
		{"MALE", 0},
		{"FEMALE", 1},
		{"OTHER", 2},
		{"HELLO", -1},
		{"NONE", -1},
	}
	for _, test := range tests {
		var model Gender
		gender, _ := model.GetByName(test.name)
		assert.Equal(gender, test.expectedResult)
	}
}
