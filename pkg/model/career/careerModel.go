package career

type Career struct {
	Id          uint64 `json:"id"`
	Acronym     string `json:"acronym"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	IsVerified  bool   `json:"is_verified"`
	DocVerifier string `json:"doc_verifier"`
	Reference   string `json:"reference"`
	FacultyId   uint64 `json:"faculty_id"`
	AlbumId     uint64 `json:"album_id"`
}
