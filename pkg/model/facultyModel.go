package model

import "time"

type Faculty struct {
	Id               uint64    `json:"id"`
	Acronym          string    `json:"acronym"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Logo             []byte    `json:"logo"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	UniversityId     uint64    `json:"university_id"`
	CreatedBy        uint64    `json:"created_by"`
}
