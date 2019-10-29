package common

import (
	"errors"
	"fmt"
)

// for enums define a type of the enum that one want
type InstanceStatus int8

// we use iota to provide int values to each of the contant types
const (
	UNVERIFIED InstanceStatus = 1 + iota
	PENDING
	VERIFIED
)

// Enums should be able to printout as strings
var InstanceStatusString = [...]string{
	"UNVERIFIED",
	"PENDING",
	"VERIFIED",
}

// String() function will return the english name
func (instanceStatus InstanceStatus) String() string {
	err := instanceStatus.IsValid()
	if err == nil {
		return InstanceStatusString[instanceStatus-1]
	} else {
		return ""
	}
}

func (instanceStatus InstanceStatus) GetByName(name string) (InstanceStatus, error) {
	for i, instanceStatus := range InstanceStatusString {
		if instanceStatus == name {
			return InstanceStatus(i + 1), nil
		}
	}
	err := fmt.Sprintf("The name does not exist: %s", name)
	return 0, errors.New(err)
}

func (instanceStatus InstanceStatus) IsValid() error {
	ok := instanceStatus > 0 && instanceStatus <= InstanceStatus(len(InstanceStatusString))
	if !ok {
		err := fmt.Sprintf("Index out of the range, index: %d; range: %d", int8(instanceStatus), len(InstanceStatusString))
		return errors.New(err)
	}
	return nil
}
