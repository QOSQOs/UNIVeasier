package person

import "time"

type Person struct {
	Id            uint64    `json:"id"`
	Code          string    `json:"code"`
	Gender        int32     `json:"gender"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Avatar        []byte    `json:"avatar"`
	Type          int32     `json:"type"`
	HomeCity      string    `json:"home_city"`
	CurrentCity   string    `json:"current_city"`
	Ethnic        int32     `json:"ethnic"`
	Nationality   string    `json:"nationality"`
	BirthDate     time.Time `json:"birth_date"`
	AdmissionYear int32     `json:"admission_year"`
	Period        int32     `json:"period"`
	IsVerified    bool      `json:"is_verified"`
	DocVerifier   []byte    `json:"doc_verifier"`
}
