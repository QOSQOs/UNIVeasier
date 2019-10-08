package university

type University struct {
	Id               uint64  `json:"id"`
	LocalRank        int32   `json:"local_rank"`
	GlobalRank       int32   `json:"global_rank"`
	TypeUniversityId uint64  `json:"type_university_id"`
	Acronym          string  `json:"acronym"`
	Name             string  `json:"name"`
	Region           string  `json:"region"`
	Description      string  `json:"description"`
	Latitude         float32 `json:"latitude"`
	Altitude         float32 `json:"altitude"`
	Logo             string  `json:"logo"`
	IsVerified       bool    `json:"is_verified"`
	DocVerifier      string  `json:"doc_verifier"`
}
