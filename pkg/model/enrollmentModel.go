package model

import "time"

type Enrollement struct {
	Id               uint64    `json:"id"`
	Status           int32     `json:"status"`
	Grade            float32   `json:"grade"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	PersonId         uint64    `json:"person_id"`
	CourseId         uint64    `json:"course_id"`
}
