package model

import (
	"errors"
	"time"
)

type TypeUniversity struct {
	Id               uint64    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedBy        uint64    `json:"created_by"`
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
