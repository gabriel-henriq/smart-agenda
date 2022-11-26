package models

import "time"

type Room struct {
	ID        int32     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type RoomList struct {
	Rooms      []Room `json:"rooms"`
	Pagination `json:"pagination"`
}
