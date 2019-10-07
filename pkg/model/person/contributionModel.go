package person

type Contribution struct {
	Id       uint  `json:"id"`
	Date     int64 `json:"date"`
	Score    int64 `json:"score"`
	PersonId uint  `json:"person_id"`
}
