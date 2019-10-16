package complement

import "time"

type Award struct {
	Id               uint64    `json:"id"`
	Year             string    `json:"year"`
	Period           string    `json:"period"`
	Description      string    `json:"description"`
	Event            string    `json:"event"`
	Place            int32     `json:"place"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CareerId         uint64    `json:"career_id"`
	AlbumId          uint64    `json:"album_id"`
	CreatedBy        uint64    `json:"created_by"`
}
