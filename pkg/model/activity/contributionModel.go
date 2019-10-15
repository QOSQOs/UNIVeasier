package activity

import "time"

type Contribution struct {
	Id          uint64    `json:"id"`
	Date        time.Time `json:"date"`
	Score       int64     `json:"score"`
	Description string    `json:"description"`
	PersonId    uint64    `json:"person_id"`
}
