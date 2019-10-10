package course

type Course struct {
	Id           uint64 `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Credits      int32  `json:"credits"`
	Group        int32  `json:"group"`
	Year         int32  `json:"year"`
	Period       int32  `json:"period"`
	Semester     int32  `json:"semester"`
	TotalHours   int32  `json:"total_hours"`
	IsVerified   bool   `json:"is_verified"`
	DocVerifier  string `json:"doc_verifier"`
	TypeCourseId uint64 `json:"type_course_id"`
	CareerId     uint64 `json:"carrera_id"`
	PersonId     uint64 `json:"person_id"`
}
