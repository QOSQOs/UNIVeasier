package complement

import "time"

type Schedule struct {
	Id          uint64    `json:"id"`
	CourseId    uint64    `json:"course_id"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Day         int32     `json:"day"`
	IsVerified  bool      `json:"is_verified"`
	DocVerifier string    `json:"doc_verifier"`
}
