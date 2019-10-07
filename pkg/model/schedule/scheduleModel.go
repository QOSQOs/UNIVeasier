package schedule

type Schedule struct {
	Id         uint  `json:"id"`
	CourseId   uint  `json:"course_id"`
	Start      int64 `json:"start"`
	End        int64 `json:"end"`
	Day        uint  `json:"day"`
	IsVerified bool  `json:"is_verified"`
}
