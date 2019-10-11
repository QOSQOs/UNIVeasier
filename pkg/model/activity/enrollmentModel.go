package activity

type Enrollement struct {
	Id          uint64  `json:"id"`
	PersonId    uint64  `json:"person_id"`
	CourseId    uint64  `json:"course_id"`
	Status      int32   `json:"status"`
	Grade       float32 `json:"grade"`
	IsVerified  bool    `json:"is_verified"`
	DocVerifier []byte  `json:"doc_verifier"`
}
