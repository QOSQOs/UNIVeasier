package complement

type Interest struct {
	Id       uint64 `json:"id"`
	Tag      string `json:"tag"`
	IsSkill  bool   `json:"is_skill"`
	PersonId uint64 `json:"person_id"`
}
