package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypePerson(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		typePerson      TypePerson
		expectedResult1 string
		expectedResult2 bool
	}{
		{1, "VISITOR", true},
		{2, "STUDENT", true},
		{3, "GRADUATED", true},
		{4, "PROFESSOR", true},
		{5, "", false},
		{6, "", false},
	}
	for _, test := range tests {
		name := test.typePerson.String()
		assert.Equal(name, test.expectedResult1)
		err := test.typePerson.IsValid()
		assert.Equal((err == nil), test.expectedResult2)
	}
}

func TestTypePersonGetByName(t *testing.T) {
	InitComplements()
	assert := assert.New(t)

	var tests = []struct {
		name           string
		expectedResult TypePerson
	}{
		{"VISITOR", 1},
		{"STUDENT", 2},
		{"GRADUATED", 3},
		{"PROFESSOR", 4},
		{"HELLO", 0},
		{"NONE", 0},
	}
	for _, test := range tests {
		var model TypePerson
		typePerson, _ := model.GetByName(test.name)
		assert.Equal(typePerson, test.expectedResult)
	}
}
