package faculty

type Faculty struct {
	Id           uint   `json:"id"`
	UniversityId uint   `json:"university_id"`
	Acronym      string `json:"acronym"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Logo         string `json:"logo"`
	IsVerified   bool   `json:"is_verified"`
}
