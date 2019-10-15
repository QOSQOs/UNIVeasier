package course

type TypeCourse struct {
	Id           uint64 `json:"id"`
	Acronym      string `json:"acronym"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsVerified   bool   `json:"is_verified"`
	DocVerifier  []byte `json:"doc_verifier"`
	UniversityId uint64 `json:"university_id"`
}
