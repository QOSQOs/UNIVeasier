package complement

import "time"

type Interest struct {
	Id               uint64    `json:"id"`
	Tag              string    `json:"tag"`
	IsSkill          bool      `json:"is_skill"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedBy        uint64    `json:"created_by"`
}
