package activity

import "time"

type Comment struct {
	Id               uint64    `json:"id"`
	Content          string    `json:"content"`
	Likes            string    `json:"likes"`
	Dislikes         int64     `json:"dislikes"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	StoryId          uint64    `json:"story_id"`
	CreatedBy        uint64    `json:"created_by"`
}
