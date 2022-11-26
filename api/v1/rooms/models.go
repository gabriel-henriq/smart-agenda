package room

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"math"
)

func (r Room) toJSONRoom(sqlRoom sqlc.Room) models.Room {
	return models.Room{
		ID:        sqlRoom.ID,
		Name:      sqlRoom.Name.String,
		CreatedAt: sqlRoom.CreatedAt.Time,
		UpdatedAt: sqlRoom.UpdatedAt.Time,
	}
}

func (r Room) toJSONRoomList(SQLRooms []sqlc.ListRoomsRow, pageID, pageSize int32) models.RoomList {
	var rooms []models.Room

	for _, room := range SQLRooms {
		rooms = append(rooms, models.Room{
			Name:      room.Name.String,
			CreatedAt: room.CreatedAt.Time,
			UpdatedAt: room.UpdatedAt.Time,
		})
	}

	totalPages := int32(math.Ceil(math.Round(float64(pageID / pageSize))))

	return models.RoomList{
		Rooms: rooms,
		Pagination: models.Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalPages: totalPages,
		},
	}
}
