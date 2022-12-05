package tablet

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/middleware"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
	"github.com/gin-gonic/gin"
)

type ITablet interface {
	SetupTabletRoute(routerGroup *gin.RouterGroup)
}

type Tablet struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewTablet(db db.Store, config util.Config) ITablet {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Tablet{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (t Tablet) SetupTabletRoute(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(t.tokenMaker))
	authRoutes.POST("/tablet", t.create)
	authRoutes.GET("/tablet", t.list)
	authRoutes.PATCH("/tablet", t.update)
	authRoutes.GET("/tablet/:id", t.getByID)
	authRoutes.DELETE("/tablet/:id", t.delete)
}
