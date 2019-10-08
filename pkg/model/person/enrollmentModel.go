package person

type Enrollement struct {
	Id          uint64  `json:"id"`
	PersonId    uint64  `json:"person_id"`
	CareerId    uint64  `json:"career_id"`
	Status      int32   `json:"status"`
	Grade       float32 `json:"grade"`
	IsVerified  bool    `json:"is_verified"`
	DocVerifier string  `json:"doc_verifier"`
}
