package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type CreateRoomRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
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
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type DeleteRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type ListRoomRequest struct {
	PageSize int32  `form:"page_size"`
	PageID   int32  `form:"page_id"`
	OrderBy  string `form:"order_by"`
	Reverse  bool   `form:"reverse"`
}

type RoomList struct {
	Rooms      []RoomResponse `json:"rooms"`
	Pagination `json:"pagination"`
}

func ToJSONRoom(sqlRoom sqlc.Room) RoomResponse {
	return RoomResponse{
		ID:         sqlRoom.ID,
		Name:       sqlRoom.Name,
		LabelColor: sqlRoom.LabelColor,
		CreatedAt:  sqlRoom.CreatedAt.String(),
		UpdatedAt:  sqlRoom.UpdatedAt.String(),
	}
}

func ToJSONRoomList(SQLRooms []sqlc.ListRoomsRow, pageID, pageSize int32) RoomList {
	var rooms []RoomResponse

	for _, room := range SQLRooms {
		rooms = append(rooms, RoomResponse{
			Name:       room.Name,
			LabelColor: room.LabelColor,
			CreatedAt:  room.CreatedAt.String(),
			UpdatedAt:  room.UpdatedAt.String(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLRooms[0].TotalItems) / float64(pageSize)))

	return RoomList{
		Rooms: rooms,
		Pagination: Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLRooms[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
