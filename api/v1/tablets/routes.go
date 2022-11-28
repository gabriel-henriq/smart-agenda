package tablets

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

func (r Tablet) SetupTabletRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/tablet", r.createTablet)
	routerGroup.GET("/tablet", r.listTablet)
	routerGroup.GET("/tablet/:id", r.getTabletByID)
	routerGroup.DELETE("/tablet/:id", r.deleteTablet)
}
