package model

import "time"

type Photo struct {
	Id               uint64    `json:"id"`
	Pic              []byte    `json:"pic"`
	Description      string    `json:"description"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	AlbumId          uint64    `json:"album_id"`
	CreatedBy        uint64    `json:"created_by"`
}
