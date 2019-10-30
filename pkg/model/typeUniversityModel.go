package model

import (
	"errors"

	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
)

type TypeUniversity struct {
	Id               types.Int64           `json:"id"`
	Name             types.String          `json:"name"`
	Description      types.String          `json:"description"`
	IsVerified       common.InstanceStatus `json:"is_verified"`
	DocVerifier      []byte                `json:"doc_verifier"`
	CreatedDate      types.Time            `json:"created_date"`
	LastModifiedDate types.Time            `json:"last_modified_date"`
	CreatedBy        types.Int64           `json:"created_by"`
}

func (tu *TypeUniversity) Validate() error {
	err := tu.IsVerified.IsValid()
	if err != nil {
		return err
	}

	if !tu.CreatedBy.Valid {
		return errors.New("The CreatedBy attribute can not be NULL")
	}
	return nil
}
