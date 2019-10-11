package complement

type Photo struct {
	Id          uint64 `json:"id"`
	Pic         []byte `json:"pic"`
	Description string `json:"description"`
	AlbumId     uint64 `json:"album_id"`
}
