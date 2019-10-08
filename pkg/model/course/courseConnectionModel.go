package course

type CourseConnection struct {
	Id               uint64 `json:"id"`
	CareerId         uint64 `json:"career_id"`
	Year             int32  `json:"year"`
	Period           int32  `json:"period"`
	Child            uint64 `json:"child"`
	Parent           uint64 `json:"parent"`
	ThresholdCredits int32  `json:"threshold _credits"`
}
