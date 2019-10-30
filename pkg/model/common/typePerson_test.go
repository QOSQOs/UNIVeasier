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
		{0, "VISITOR", true},
		{1, "STUDENT", true},
		{2, "GRADUATED", true},
		{3, "PROFESSOR", true},
		{4, "", false},
		{5, "", false},
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
		{"VISITOR", 0},
		{"STUDENT", 1},
		{"GRADUATED", 2},
		{"PROFESSOR", 3},
		{"HELLO", -1},
		{"NONE", -1},
	}
	for _, test := range tests {
		var model TypePerson
		typePerson, _ := model.GetByName(test.name)
		assert.Equal(typePerson, test.expectedResult)
	}
}
