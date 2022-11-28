package models

import "time"

type Tablet struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TabletList struct {
	Tablets    []Tablet `json:"tablets"`
	Pagination `json:"pagination"`
}
