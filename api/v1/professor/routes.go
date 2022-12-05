package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/middleware"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
	"github.com/gin-gonic/gin"
)

type IProfessor interface {
	SetupProfessorRoute(routerGroup *gin.RouterGroup)
}

type Professor struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewProfessor(db db.Store, config util.Config) IProfessor {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Professor{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (p Professor) SetupProfessorRoute(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(p.tokenMaker))
	authRoutes.POST("/professor", p.create)
	authRoutes.GET("/professor", p.list)
	authRoutes.PATCH("/professor/", p.update)
	authRoutes.GET("/professor/:id", p.get)
	authRoutes.DELETE("/professor/:id", p.delete)
}
