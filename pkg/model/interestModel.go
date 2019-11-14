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

func (interestModel *Interest) Validate() error {
	if len(interestModel.Tag.String) == 0 || len(interestModel.Tag.String) > 100 {
		return &errors.InvalidNameError{interestModel.Tag.String}
	}

	if interestModel.CreatedBy.IsNull() {
		return &errors.NullValueNotAllowedError{"CreatedBy"}
	}

	return nil
}
