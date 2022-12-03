package room

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type createRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type updateRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type response struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type deleteRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type list struct {
	Rooms                 []response `json:"rooms"`
	v1.PaginationResponse `json:"pagination"`
}

func toJSON(sqlRoom sqlc.Room) response {
	return response{
		ID:         sqlRoom.ID,
		Name:       sqlRoom.Name,
		LabelColor: sqlRoom.LabelColor,
		CreatedAt:  sqlRoom.CreatedAt.String(),
		UpdatedAt:  sqlRoom.UpdatedAt.String(),
	}
}

func toJSONList(SQLRooms []sqlc.ListRoomsRow, pageID, pageSize int32) list {
	var rooms []response

	for _, room := range SQLRooms {
		rooms = append(rooms, response{
			ID:         room.ID,
			Name:       room.Name,
			LabelColor: room.LabelColor,
			CreatedAt:  room.CreatedAt.String(),
			UpdatedAt:  room.UpdatedAt.String(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLRooms[0].TotalItems) / float64(pageSize)))

	return list{
		Rooms: rooms,
		PaginationResponse: v1.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLRooms[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
