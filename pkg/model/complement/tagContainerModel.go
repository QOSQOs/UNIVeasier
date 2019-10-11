package complement

type TagContainer struct {
	Id        uint64 `json:"id"`
	Tag       string `json:"tag"`
	Agree     int64  `json:"agree"`
	Disagree  int64  `json:"disagree"`
	Reference string `json:"reference"`
	Type      int32  `json:"type"`
	Layer     int32  `json:"layer"`
	CareerId  string `json:"career_id"`
}
