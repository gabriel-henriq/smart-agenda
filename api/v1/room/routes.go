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
	routerGroup.POST("/room", r.create)
	routerGroup.GET("/room", r.list)
	routerGroup.PATCH("/room", r.update)
	routerGroup.GET("/room/:id", r.getByID)
	routerGroup.DELETE("/room/:id", r.delete)
}
