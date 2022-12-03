package tablet

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type ITablet interface {
	SetupTabletRoute(routerGroup *gin.RouterGroup)
}

type Tablet struct {
	db db.Store
}

func NewTablet(db db.Store) ITablet {
	return Tablet{
		db: db,
	}
}

func (t Tablet) SetupTabletRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/tablet", t.create)
	routerGroup.GET("/tablet", t.list)
	routerGroup.PATCH("/tablet", t.update)
	routerGroup.GET("/tablet/:id", t.getByID)
	routerGroup.DELETE("/tablet/:id", t.delete)
}
