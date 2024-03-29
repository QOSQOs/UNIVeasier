package model

import "time"

type Career struct {
	Id               uint64    `json:"id"`
	Acronym          string    `json:"acronym"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Logo             []byte    `json:"logo"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	Reference        string    `json:"reference"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	FacultyId        uint64    `json:"faculty_id"`
	AlbumId          uint64    `json:"album_id"`
	CreatedBy        uint64    `json:"created_by"`
}
