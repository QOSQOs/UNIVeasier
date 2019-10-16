package person

import "time"

type CareerPerson struct {
	CareerId         uint64    `json:"career_id"`
	PersonId         uint64    `json:"person_id"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
}
