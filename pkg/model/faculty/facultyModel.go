package faculty

type Faculty struct {
	Id           uint64 `json:"id"`
	UniversityId uint64 `json:"university_id"`
	Acronym      string `json:"acronym"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Logo         string `json:"logo"`
	IsVerified   bool   `json:"is_verified"`
	DocVerifier  string `json:"doc_verifier"`
}
