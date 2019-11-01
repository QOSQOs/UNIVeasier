package errors

import (
	"fmt"
)

type InvalidGenderError struct {
	Gender string
}

func (e *InvalidGenderError) Error() string {
	return fmt.Sprintf("Invalid gender type value: %q", e.Gender)
}

type InvalidGenderIndiceError struct {
	GenderIndice int8
}

func (e *InvalidGenderIndiceError) Error() string {
	return fmt.Sprintf("Could not find a gender type associated with the index %d", e.GenderIndice)
}

type InvalidStatusError struct {
	Status string
}

func (e *InvalidStatusError) Error() string {
	return fmt.Sprintf("Invalid status type value: %q", e.Status)
}

type InvalidStatusIndiceError struct {
	StatusIndice int8
}

func (e *InvalidStatusIndiceError) Error() string {
	return fmt.Sprintf("Could not find a status type associated with the index %d", e.StatusIndice)
}

type InvalidPersonError struct {
	Person string
}

func (e *InvalidPersonError) Error() string {
	return fmt.Sprintf("Invalid person type value: %q", e.Person)
}

type InvalidPersonIndiceError struct {
	PersonIndice int8
}

func (e *InvalidPersonIndiceError) Error() string {
	return fmt.Sprintf("Could not find a person type associated with the index %d", e.PersonIndice)
}

type OverflowError struct {
	Type string
}

func (e *OverflowError) Error() string {
	return fmt.Sprintf("Overflow error converting value to data type %q", e.Type)
}

type InvalidTypeError struct {
	ExpectedType string
	Type         string
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("Cannot convert %q to %q", e.Type, e.ExpectedType)
}
