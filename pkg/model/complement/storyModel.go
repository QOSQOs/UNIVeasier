package complement

import "time"

type Story struct {
	Id               uint64    `json:"id"`
	Title            string    `json:"title"`
	Body             string    `json:"body"`
	Recommended      int64     `json:"recomended"`
	Unrecommended    int64     `json:"unrecomended"`
	Views            int64     `json:"views"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	AllowComments    bool      `json:"allow_comments"`
	Layer            int32     `json:"layer"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedBy        uint64    `json:"created_by"`
	CareerId         uint64    `json:"career_id"`
}
