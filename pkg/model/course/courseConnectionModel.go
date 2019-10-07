package course

type CourseConnection struct {
	Id               uint `json:"id"`
	CareerId         uint `json:"career_id"`
	Year             int  `json:"year"`
	Period           int  `json:"period"`
	Child            uint `json:"child"`
	Parent           uint `json:"parent"`
	ThresholdCredits int  `json:"threshold _credits"`
}
