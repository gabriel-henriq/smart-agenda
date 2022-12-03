package professor

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
	routerGroup.POST("/professor", p.create)
	routerGroup.GET("/professor", p.list)
	routerGroup.PATCH("/professor/", p.update)
	routerGroup.GET("/professor/:id", p.get)
	routerGroup.DELETE("/professor/:id", p.delete)
}
