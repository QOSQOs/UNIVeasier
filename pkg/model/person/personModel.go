package person

type Person struct {
	Id            uint   `json:"id"`
	Code          string `json:"code"`
	Gender        int    `json:"gender"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	Avatar        string `json:"avatar"`
	Type          int    `json:"type"`
	HomeCity      string `json:"home_city"`
	CurrentCity   string `json:"current_city"`
	Ethnic        int    `json:"ethnic"`
	Nationality   string `json:"nationality"`
	BirthDate     int64  `json:"birth_date"`
	AdmissionYear int    `json:"admission_year"`
	Period        int    `json:"period"`
	IsVerified    bool   `json:"is_verified"`
}
