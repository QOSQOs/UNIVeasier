package test

import (
	"errors"
	"fmt"

	"github.com/QOSQOs/UNIVeasier/internal/common"
)

type Test struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (t *Test) Validate() error {
	ok := common.Letters.MatchString(t.Name)
	if len(t.Name) == 0 || !ok {
		return errors.New("Invalid test name")
	}

	fmt.Println("2")

	if t.Age < 0 || t.Age > 200 {
		return errors.New("Invalid test age")
	}

	fmt.Println("3")

	ok = common.Emails.MatchString(t.Email)
	if len(t.Email) == 0 || !ok {
		return errors.New("Invalid test email")
	}
	fmt.Println("4")
	return nil
}
