package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
)

type Person int8

const (
	UNKNOWN_PERSON Person = -1 + iota
	VISITOR
	STUDENT
	GRADUATED
	PROFESSOR
)

var personList = [...]string{
	"VISITOR",
	"STUDENT",
	"GRADUATED",
	"PROFESSOR",
}

func (p Person) ToString() (string, error) {
	if err := p.IsValid(); err == nil {
		return personList[p], nil
	} else {
		return "", err
	}
}

func (p Person) IsValid() error {
	if p < 0 || int(p) >= len(personList) {
		return &errors.InvalidPersonIndiceError{int8(p)}
	}
	return nil
}

func ToPerson(person string) (Person, error) {
	for i, val := range personList {
		if val == person {
			return Person(i), nil
		}
	}
	return UNKNOWN_PERSON, &errors.InvalidPersonError{person}
}
