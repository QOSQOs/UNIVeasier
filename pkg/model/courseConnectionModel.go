package model

import "time"

type CoursesConnection struct {
	Id               uint64    `json:"id"`
	Year             int32     `json:"year"`
	Period           int32     `json:"period"`
	ThresholdCredits int32     `json:"threshold _credits"`
	ChildCourse      uint64    `json:"child_course"`
	ParentCourse     uint64    `json:"parent_course"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CareerId         uint64    `json:"career_id"`
	CreatedBy        uint64    `json:"created_by"`
}
