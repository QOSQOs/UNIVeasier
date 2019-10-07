package person

type Enrollement struct {
	Id         uint    `json:"id"`
	PersonId   uint    `json:"person_id"`
	CareerId   uint    `json:"career_id"`
	Status     uint    `json:"status"`
	Grade      float32 `json:"grade"`
	IsVerified bool    `json:"is_verified"`
}
