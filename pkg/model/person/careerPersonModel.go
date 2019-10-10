package person

type CareerPerson struct {
	PersonId    uint64 `json:"person_id"`
	CareerId    uint64 `json:"career_id"`
	IsVerified  bool   `json:"is_verified"`
	DocVerifier string `json:"doc_verifier"`
}
