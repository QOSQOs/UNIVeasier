package activity

import "time"

type Contribution struct {
	Id               uint64    `json:"id"`
	Score            int64     `json:"score"`
	Description      string    `json:"description"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	PersonId         uint64    `json:"person_id"`
}
