package university

type TypeUniversity struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsVerified  bool   `json:"is_verified"`
	DocVerifier string `json:"doc_verifier"`
}
