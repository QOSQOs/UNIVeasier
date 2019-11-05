package model

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
)

type Album struct {
	Id               types.Int64  `json:"id"`
	Name             types.String `json:"name"`
	Description      types.String `json:"description"`
	CreatedDate      types.Time   `json:"created_date"`
	LastModifiedDate types.Time   `json:"last_modified_date"`
	CreatedBy        types.Int64  `json:"created_by"`
}

func (albumModel *Album) Validate() error {

	if len(albumModel.Name.String) == 0 || len(albumModel.Name.String) > 100 {
		return &errors.InvalidNameError{albumModel.Name.String}
	}

	if len(albumModel.Description.String) == 0 {
		return &errors.InvalidNameError{albumModel.Description.String}
	}

	if albumModel.CreatedBy.IsNull() {
		return &errors.NullValueNotAllowedError{"CreatedBy"}
	}

	return nil
}
