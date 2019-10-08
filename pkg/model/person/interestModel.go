package person

type Interest struct {
	Id       uint64 `json:"id"`
	Tag      string `json:"tag"`
	Skill    bool   `json:"skill"`
	PersonId uint64 `json:"person_id"`
}
