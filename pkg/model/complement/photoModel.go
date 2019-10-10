package complement

type Photo struct {
	Id          uint64 `json:"id"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	AlbumId     uint64 `json:"album_id"`
}
