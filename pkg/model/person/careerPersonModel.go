package person

type CareerPerson struct {
	Id         uint `json:"id"`
	PersonId   uint `json:"person_id"`
	CareerId   uint `json:"career_id"`
	IsVerified bool `json:"is_verified"`
}
