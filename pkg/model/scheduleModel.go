package model

import "time"

type Schedule struct {
	Id               uint64    `json:"id"`
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
	Day              int32     `json:"day"`
	Room             string    `json:"room"`
	TypeLesson       int32     `json:"type_lesson"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CourseId         uint64    `json:"course_id"`
	CreatedBy        uint64    `json:"created_by"`
}
