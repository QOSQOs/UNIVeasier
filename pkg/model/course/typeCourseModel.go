package course

type TypeCourse struct {
	Id          uint   `json:"id"`
	Acronym     string `json:"acronym"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsVerified  bool   `json:"is_verified"`
}
