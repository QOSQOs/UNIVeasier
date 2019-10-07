package career

type Career struct {
	Id          uint   `json:"id"`
	FacultyId   uint   `json:"faculty_id"`
	Acronym     string `json:"acronym"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	IsVerified  bool   `json:"is_verified"`
}
