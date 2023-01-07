package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

type CreateRoomRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type DeleteRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type UpdateRoomRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type RoomResponse struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type ListRoomResponse struct {
	Rooms              []RoomResponse `json:"rooms"`
	PaginationResponse `json:"pagination"`
}

func RoomToJSON(sqlRoom sqlc.Room) RoomResponse {
	return RoomResponse{
		ID:         sqlRoom.ID,
		Name:       sqlRoom.Name,
		LabelColor: sqlRoom.LabelColor,
		CreatedAt:  sqlRoom.CreatedAt.Unix(),
		UpdatedAt:  sqlRoom.UpdatedAt.Unix(),
	}
}

func RoomsToJSONList(SQLRooms []sqlc.ListRoomsRow) ListRoomResponse {
	var rooms []RoomResponse

	for _, room := range SQLRooms {
		rooms = append(rooms, RoomResponse{
			ID:         room.ID,
			Name:       room.Name,
			LabelColor: room.LabelColor,
			CreatedAt:  room.CreatedAt.Unix(),
			UpdatedAt:  room.UpdatedAt.Unix(),
		})
	}

	return ListRoomResponse{Rooms: rooms}
}
