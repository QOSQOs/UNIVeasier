package university

import "time"

type University struct {
	Id               uint64    `json:"id"`
	LocalRank        int32     `json:"local_rank"`
	GlobalRank       int32     `json:"global_rank"`
	Acronym          string    `json:"acronym"`
	Name             string    `json:"name"`
	Region           string    `json:"region"`
	Description      string    `json:"description"`
	Latitude         float32   `json:"latitude"`
	Longitude        float32   `json:"longitude"`
	Logo             []byte    `json:"logo"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	Reference        string    `json:"reference"`
	TypeUniversityId uint64    `json:"type_university_id"`
	AlbumId          uint64    `json:"album_id"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedBy        uint64    `json:"created_by"`
}
