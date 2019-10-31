package model

import "time"

type Course struct {
	Id               uint64    `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Credits          int32     `json:"credits"`
	Batch            int32     `json:"batch"`
	Year             int32     `json:"year"`
	Period           int32     `json:"period"`
	Semester         int32     `json:"semester"`
	TotalHours       int32     `json:"total_hours"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	TypeCourseId     uint64    `json:"type_course_id"`
	CareerId         uint64    `json:"carrera_id"`
	ProfessorId      uint64    `json:"professor_id"`
	CreatedBy        uint64    `json:"created_by"`
}
