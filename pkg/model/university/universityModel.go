package university

type University struct {
	Id               uint    `json:"id"`
	LocalRank        int     `json:"local_rank"`
	GlobalRank       int     `json:"global_rank"`
	TypeUniversityId uint    `json:"type_university_id"`
	Acronym          string  `json:"acronym"`
	Name             string  `json:"name"`
	Region           string  `json:"region"`
	Description      string  `json:"description"`
	Latitude         float32 `json:"latitude"`
	Altitude         float32 `json:"altitude"`
	Logo             string  `json:"logo"`
	IsVerified       bool    `json:"is_verified"`
}
