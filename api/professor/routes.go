package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type IProfessor interface {
	SetupProfessorRoute(router *gin.Engine)
}

type Professor struct {
	db.Store
}

func NewProfessor(db db.Store) IProfessor {
	return Professor{
		Store: db,
	}
}

func (p Professor) SetupProfessorRoute(router *gin.Engine) {

	router.POST("/professor", p.createProfessor)
	router.GET("/professor", p.listProfessor)
	router.GET("/professor/:id", p.getProfessor)
	router.DELETE("/professor/:id", p.deleteProfessor)

}
