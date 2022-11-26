package room

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IRoom interface {
	SetupRoomRoute(router *gin.Engine)
}

type Room struct {
	db db.Store
}

func NewRoom(db db.Store) IRoom {
	return Room{
		db: db,
	}
}

func (r Room) SetupRoomRoute(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/room", r.createRooms)
		v1.GET("/room", r.listRoom)
		v1.GET("/room/:id", r.getRoomByID)
		v1.DELETE("/room/:id", r.deleteRoom)
	}
}
