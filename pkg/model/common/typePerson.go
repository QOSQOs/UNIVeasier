package common

import (
	"errors"
	"fmt"
)

// for enums define a type of the enum that one want
type TypePerson int8

// we use iota to provide int values to each of the contant types
const (
	VISITOR TypePerson = 0 + iota
	STUDENT
	GRADUATED
	PROFESSOR
)

// Enums should be able to printout as strings
var TypePersonString = [...]string{
	"VISITOR",
	"STUDENT",
	"GRADUATED",
	"PROFESSOR",
}

// String() function will return the english name
func (typePerson TypePerson) String() string {
	err := typePerson.IsValid()
	if err == nil {
		return TypePersonString[typePerson]
	} else {
		return ""
	}
}

func (typePerson TypePerson) GetByName(name string) (TypePerson, error) {
	for i, typePerson := range TypePersonString {
		if typePerson == name {
			return TypePerson(i), nil
		}
	}
	err := fmt.Sprintf("The name does not exist: %s", name)
	return -1, errors.New(err)
}

func (typePerson TypePerson) IsValid() error {
	ok := typePerson >= 0 && typePerson < TypePerson(len(TypePersonString))
	if !ok {
		err := fmt.Sprintf("TypePerson: Index out of the range, index: %d; range: %d", int8(typePerson), len(TypePersonString))
		return errors.New(err)
	}
	return nil
}
