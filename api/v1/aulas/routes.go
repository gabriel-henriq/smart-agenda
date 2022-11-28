package aulas

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

func (r Aula) SetupAulaRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/aula", r.createAula)
	routerGroup.GET("/aula", r.listAula)
	routerGroup.GET("/aula/:id", r.getAulaByID)
	routerGroup.DELETE("/aula/:id", r.deleteAula)
}
