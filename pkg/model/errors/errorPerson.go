package errors

import (
	"fmt"
)

type InvalidNameError struct {
	Name string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("Invalid name: %q", e.Name)
}

type InvalidLastNameError struct {
	LastName string
}

func (e *InvalidLastNameError) Error() string {
	return fmt.Sprintf("Invalid last name: %q", e.LastName)
}

type InvalidEmailError struct {
	Email string
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("Invalid email: %q", e.Email)
}
