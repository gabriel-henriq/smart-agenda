package room

import (
	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/api/v1/middleware"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
)

type IRoom interface {
	SetupRoomRoute(routerGroup *gin.RouterGroup)
}

type Room struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewRoom(db db.Store, config util.Config) IRoom {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Room{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (r Room) SetupRoomRoute(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(r.tokenMaker))
	authRoutes.POST("/room", r.create)
	authRoutes.GET("/room", r.list)
	authRoutes.PATCH("/room", r.update)
	authRoutes.GET("/room/:id", r.getByID)
	authRoutes.DELETE("/room/:id", r.delete)
}
