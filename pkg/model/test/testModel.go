package test

import (
	"errors"
	"regexp"
)

var lettersRegex = regexp.MustCompile(`^[a-zA-Z]+\s*[a-zA-Z]*$`)
var emailsRegex = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z.\-]+\.[a-z]{2,4}$`)

type Test struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (t *Test) Validate() error {
	ok := lettersRegex.MatchString(t.Name)
	if len(t.Name) == 0 || !ok {
		return errors.New("Invalid test name")
	}

	if t.Age < 0 || t.Age > 200 {
		return errors.New("Invalid test age")
	}

	ok = emailsRegex.MatchString(t.Email)
	if len(t.Email) == 0 || !ok {
		return errors.New("Invalid test email")
	}

	return nil
}
