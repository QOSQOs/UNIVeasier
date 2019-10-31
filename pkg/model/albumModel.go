package model

import "time"

type Album struct {
	Id               uint64    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CreatedDate      time.Time `json:"created_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedBy        uint64    `json:"created_by"`
}
