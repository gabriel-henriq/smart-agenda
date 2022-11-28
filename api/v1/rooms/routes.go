package room

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IRoom interface {
	SetupRoomRoute(routerGroup *gin.RouterGroup)
}

type Room struct {
	db db.Store
}

func NewRoom(db db.Store) IRoom {
	return Room{
		db: db,
	}
}

func (r Room) SetupRoomRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/room", r.createRoom)
	routerGroup.GET("/room", r.listRoom)
	routerGroup.GET("/room/:id", r.getRoomByID)
	routerGroup.DELETE("/room/:id", r.deleteRoom)
}
