package complement

type Award struct {
	Id          uint64 `json:"id"`
	Year        string `json:"year"`
	Period      string `json:"period"`
	Description string `json:"description"`
	Event       string `json:"event"`
	place       bool   `json:"place"`
	IsVerified  bool   `json:"is_verified"`
	DocVerifier []byte `json:"doc_verifier"`
	CareerId    string `json:"career_id"`
	AlbumId     uint64 `json:"album_id"`
}
