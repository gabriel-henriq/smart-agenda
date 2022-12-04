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

type deleteRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
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
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type listResponse struct {
	Rooms                 []response `json:"rooms"`
	v1.PaginationResponse `json:"pagination"`
}

func toJSON(sqlRoom sqlc.Room) response {
	return response{
		ID:         sqlRoom.ID,
		Name:       sqlRoom.Name,
		LabelColor: sqlRoom.LabelColor,
		CreatedAt:  sqlRoom.CreatedAt.Unix(),
		UpdatedAt:  sqlRoom.UpdatedAt.Unix(),
	}
}

func toJSONList(SQLRooms []sqlc.ListRoomsRow, pageID, pageSize int32) listResponse {
	var rooms []response

	for _, room := range SQLRooms {
		rooms = append(rooms, response{
			ID:         room.ID,
			Name:       room.Name,
			LabelColor: room.LabelColor,
			CreatedAt:  room.CreatedAt.Unix(),
			UpdatedAt:  room.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLRooms[0].TotalItems) / float64(pageSize)))

	return listResponse{
		Rooms: rooms,
		PaginationResponse: v1.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLRooms[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
