package course

type Course struct {
	Id           uint   `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CarreraId    uint   `json:"carrera_id"`
	Credits      int    `json:"credits"`
	Group        int    `json:"group"`
	Year         int    `json:"year"`
	Period       int    `json:"period"`
	Semester     int    `json:"semester"`
	PersonId     uint   `json:"person_id"`
	TotalHours   int    `json:"total_hours"`
	TypeCourseId uint   `json:"type_course_id"`
	IsVerified   bool   `json:"is_verified"`
}
