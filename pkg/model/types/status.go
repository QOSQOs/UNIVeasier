package types

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
)

type Status int8

const (
	UNKNOWN_STATUS Status = -1 + iota
	UNVERIFIED
	PENDING
	VERIFIED
)

var statusList = [...]string{
	"UNVERIFIED",
	"PENDING",
	"VERIFIED",
}

func (s Status) ToString() (string, error) {
	if err := s.IsValid(); err == nil {
		return statusList[s], nil
	} else {
		return "", err
	}
}

func (s Status) IsValid() error {
	if s < 0 || int(s) >= len(statusList) {
		return &errors.InvalidStatusIndiceError{int8(s)}
	}
	return nil
}

func ToStatus(status string) (Status, error) {
	for i, val := range statusList {
		if val == status {
			return Status(i), nil
		}
	}
	return UNKNOWN_STATUS, &errors.InvalidStatusError{status}
}
