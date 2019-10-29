package model

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
		{"juan gabriel jose", 27, "juan10@qosqo.com", false},
		{"pepe", 28, "pepe@univeasier.net", false},
		{"", 12, "a@pe.edu", true},
		{"luis10", 16, "a@pe.edu", true},
		{"juan", -1, "a@pe.edu", true},
		{"gabriel", 18, "gabi@pe.edu1", true},
	}

	for _, t := range tests {
		model := Test{Name: t.name, Age: t.age, Email: t.email}
		err := model.Validate()
		if t.expectedError {
			assert.Error(err, model)
		} else {
			assert.NoError(err, model)
		}
	}
}
