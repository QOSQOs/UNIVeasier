package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
)

type Gender int8

const (
	UNKNOWN_GENDER Gender = -1 + iota
	MALE
	FEMALE
	OTHER
)

var genderList = [...]string{
	"MALE",
	"FEMALE",
	"OTHER",
}

func (g Gender) ToString() (string, error) {
	if err := g.IsValid(); err == nil {
		return genderList[g], nil
	} else {
		return "", err
	}
}

func (g Gender) IsValid() error {
	if g < 0 || int(g) >= len(genderList) {
		return &errors.InvalidGenderIndiceError{int8(g)}
	}
	return nil
}

func ToGender(gender string) (Gender, error) {
	for i, val := range genderList {
		if val == gender {
			return Gender(i), nil
		}
	}
	return UNKNOWN_GENDER, &errors.InvalidGenderError{gender}
}
