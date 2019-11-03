package model

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
)

type TypeUniversity struct {
	Id               types.Int64  `json:"id"`
	Name             types.String `json:"name"`
	Description      types.String `json:"description"`
	IsVerified       types.Status `json:"is_verified"`
	DocVerifier      []byte       `json:"doc_verifier"`
	CreatedDate      types.Time   `json:"created_date"`
	LastModifiedDate types.Time   `json:"last_modified_date"`
	CreatedBy        types.Int64  `json:"created_by"`
}

func (m *TypeUniversity) Validate() error {
	err := m.IsVerified.IsValid()
	if err != nil {
		return err
	}

	if m.CreatedBy.IsNull() {
		return &errors.NullValueNotAllowedError{"CreatedBy"}
	}
	return nil
}
