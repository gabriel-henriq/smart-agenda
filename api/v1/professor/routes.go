package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IProfessor interface {
	SetupProfessorRoute(router *gin.Engine)
}

type Professor struct {
	db db.Store
}

func NewProfessor(db db.Store) IProfessor {
	return Professor{
		db: db,
	}
}

func (p Professor) SetupProfessorRoute(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/professor", p.createProfessor)
		v1.GET("/professor", p.listProfessor)
		v1.GET("/professor/:id", p.getProfessor)
		v1.DELETE("/professor/:id", p.deleteProfessor)
	}
}
