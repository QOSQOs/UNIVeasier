package career

type Career struct {
	Id          uint64 `json:"id"`
	FacultyId   uint64 `json:"faculty_id"`
	Acronym     string `json:"acronym"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	IsVerified  bool   `json:"is_verified"`
	DocVerifier string `json:"doc_verifier"`
}
