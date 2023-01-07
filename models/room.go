package models

import (
	"math"

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

type ResponseRoom struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type ListRoomResponse struct {
	Rooms              []ResponseRoom `json:"rooms"`
	PaginationResponse `json:"pagination"`
}

func RoomToJSON(sqlRoom sqlc.Room) ResponseRoom {
	return ResponseRoom{
		ID:         sqlRoom.ID,
		Name:       sqlRoom.Name,
		LabelColor: sqlRoom.LabelColor,
		CreatedAt:  sqlRoom.CreatedAt.Unix(),
		UpdatedAt:  sqlRoom.UpdatedAt.Unix(),
	}
}

func RoomsToJSONList(SQLRooms []sqlc.ListRoomsRow, pageID, pageSize int32) ListRoomResponse {
	var rooms []ResponseRoom

	for _, room := range SQLRooms {
		rooms = append(rooms, ResponseRoom{
			ID:         room.ID,
			Name:       room.Name,
			LabelColor: room.LabelColor,
			CreatedAt:  room.CreatedAt.Unix(),
			UpdatedAt:  room.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLRooms[0].TotalItems) / float64(pageSize)))

	return ListRoomResponse{
		Rooms: rooms,
		PaginationResponse: PaginationResponse{
			PageSize:    pageID,
			CurrentPage: pageSize,
			TotalItems:  SQLRooms[0].TotalItems,
			TotalPages:  totalPages,
		},
	}
}
