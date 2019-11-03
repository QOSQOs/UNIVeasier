package model

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
)

type Interest struct {
	Id               types.Int64  `json:"id"`
	Tag              types.String `json:"tag"`
	IsSkill          bool         `json:"is_skill"`
	CreatedDate      types.Time   `json:"created_date"`
	LastModifiedDate types.Time   `json:"last_modified_date"`
	CreatedBy        types.Int64  `json:"created_by"`
}

func (i *Interest) Validate() error {
	ok := lettersRegex.MatchString(i.Tag.String)

	sizeTag := len(i.Tag.String)
	if sizeTag == 0 || sizeTag > 100 || !ok {
		return &errors.InvalidNameError{i.Tag.String}
	}

	if i.CreatedBy.IsNull() {
		return &errors.NullValueNotAllowedError{"CreatedBy"}
	}

	return nil
}
