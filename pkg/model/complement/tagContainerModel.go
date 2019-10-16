package complement

import "time"

type TagContainer struct {
	Id               uint64    `json:"id"`
	Tag              string    `json:"tag"`
	Agree            int64     `json:"agree"`
	Disagree         int64     `json:"disagree"`
	Reference        string    `json:"reference"`
	Type             int32     `json:"type"`
	Layer            int32     `json:"layer"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	IsVerified       int32     `json:"is_verified"`
	DocVerifier      []byte    `json:"doc_verifier"`
	CareerId         uint64    `json:"career_id"`
	CreatedBy        uint64    `json:"created_by"`
}
