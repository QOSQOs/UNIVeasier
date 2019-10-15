package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation1(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		name          string
		age           int
		email         string
		expectedError bool
	}{
		{"", 12, "aaa@aa.dd", true},
		{"aaa", 12, "aaa@aa.dd", false},
		{"aaa", 12, "a$aa@aa.dd", true},
		{"aaa", -1, "aaa@aa.dd", true},
	}

	for _, t := range tests {
		model := Test{Name: t.name, Age: t.age, Email: t.email}
		err := model.Validate()
		if t.expectedError {
			assert.Error(err)
		} else {
			assert.NoError(err)
		}
	}
}
