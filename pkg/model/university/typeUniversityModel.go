package university

type TypeUniversity struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsVerified  bool   `json:"is_verified"`
}
