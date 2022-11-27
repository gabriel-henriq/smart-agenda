package models

import "time"

type Room struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LabelColor string    `json:"labelColor"`
}

type RoomList struct {
	Rooms      []Room `json:"rooms"`
	Pagination `json:"pagination"`
}
