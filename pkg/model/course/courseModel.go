package course

type Course struct {
	Id           uint64 `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CareerId     uint64 `json:"carrera_id"`
	Credits      int32  `json:"credits"`
	Group        int32  `json:"group"`
	Year         int32  `json:"year"`
	Period       int32  `json:"period"`
	Semester     int32  `json:"semester"`
	PersonId     uint64 `json:"person_id"`
	TotalHours   int32  `json:"total_hours"`
	TypeCourseId uint64 `json:"type_course_id"`
	IsVerified   bool   `json:"is_verified"`
	DocVerifier  string `json:"doc_verifier"`
}
