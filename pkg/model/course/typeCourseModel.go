package course

import "time"

type TypeCourse struct {
	Id               uint64    `json:"id"`
	Acronym          string    `json:"acronym"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	RequiredCredits  int64     `json:"required_credits"`
	CareerId         uint64    `json:"career_id"`
	CreatedBy        uint64    `json:"created_by"`
}
