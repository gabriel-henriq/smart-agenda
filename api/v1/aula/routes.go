package aula

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IAula interface {
	SetupAulaRoute(routerGroup *gin.RouterGroup)
}

type Aula struct {
	db db.Store
}

func NewAula(db db.Store) IAula {
	return Aula{
		db: db,
	}
}

func (a Aula) SetupAulaRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/aula", a.createAula)
	routerGroup.GET("/aula", a.listAula)
	routerGroup.GET("/aula/:id", a.getAulaByID)
	routerGroup.DELETE("/aula/:id", a.deleteAula)
}
