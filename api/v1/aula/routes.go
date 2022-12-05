package aula

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/middleware"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
	"github.com/gin-gonic/gin"
)

type IAula interface {
	SetupAulaRoute(routerGroup *gin.RouterGroup)
}

type Aula struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewAula(db db.Store, config util.Config) IAula {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Aula{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (a Aula) SetupAulaRoute(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(a.tokenMaker))
	authRoutes.POST("/aula", a.create)
	authRoutes.GET("/aula", a.list)
	authRoutes.PATCH("/aula", a.update)
	authRoutes.GET("/aula/:id", a.getByID)
	authRoutes.DELETE("/aula/:id", a.delete)
}
