package activity

type Comment struct {
	Id       uint64 `json:"id"`
	Content  string `json:"content"`
	Like     string `json:"like"`
	Dislike  int64  `json:"dislike"`
	StoryId  int64  `json:"story_id"`
	PersonId uint64 `json:"person_id"`
}
