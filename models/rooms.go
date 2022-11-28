package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

type RoomList struct {
	Rooms      []sqlc.Room `json:"rooms"`
	Pagination `json:"pagination"`
}
