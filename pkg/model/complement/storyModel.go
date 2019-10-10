package complement

type Story struct {
	Id            uint64 `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Recommended   int64  `json:"recomended"`
	Unrecommended int64  `json:"unrecomended"`
	Views         int64  `json:"views"`
	IsVerified    bool   `json:"is_verified"`
	DocVerifier   []byte `json:"doc_verifier"`
	AllowComments bool   `json:"allow_comments"`
	Layer         int32  `json:"layer"`
	PersonId      uint64 `json:"person_id"`
}
