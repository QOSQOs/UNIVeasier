package common

import (
	"errors"
	"fmt"
)

// for enums define a  of the enum that one want
type Gender int8

// we use iota to provide int values to each of the contant s
const (
	MALE Gender = 0 + iota
	FEMALE
	NEUTRAL
)

// Enums should be able to printout as strings
var GenderString = [...]string{
	"MALE",
	"FEMALE",
	"OTHER",
}

// String() function will return the english name
func (gender Gender) String() string {
	err := gender.IsValid()
	if err == nil {
		return GenderString[gender]
	} else {
		return ""
	}
}

func (gender Gender) GetByName(name string) (Gender, error) {
	for i, gender := range GenderString {
		if gender == name {
			return Gender(i), nil
		}
	}
	err := fmt.Sprintf("The name does not exist: %s", name)
	return -1, errors.New(err)
}

func (gender Gender) IsValid() error {
	ok := gender >= 0 && gender < Gender(len(GenderString))
	if !ok {
		err := fmt.Sprintf("Index out of the range, index: %d; range: %d", int8(gender), len(GenderString))
		return errors.New(err)
	}
	return nil
}
