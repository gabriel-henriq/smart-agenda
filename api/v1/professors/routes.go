package professors

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IProfessor interface {
	SetupProfessorRoute(routerGroup *gin.RouterGroup)
}

type Professor struct {
	db db.Store
}

func NewProfessor(db db.Store) IProfessor {
	return Professor{
		db: db,
	}
}

func (p Professor) SetupProfessorRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/professor", p.createProfessor)
	routerGroup.GET("/professor", p.listProfessor)
	routerGroup.GET("/professor/:id", p.getProfessor)
	routerGroup.DELETE("/professor/:id", p.deleteProfessor)

}
