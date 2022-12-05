package user

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/middleware"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	SetupUserRoute(routerGroup *gin.RouterGroup)
}

type User struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewUser(db db.Store, config util.Config) IUser {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return User{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (u User) SetupUserRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/user", u.create)
	routerGroup.POST("/user/login", u.loginUser)

	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(u.tokenMaker))
	authRoutes.GET("/user", u.list)
	authRoutes.PATCH("/user", u.update)
	authRoutes.GET("/user/:id", u.getByID)
	authRoutes.DELETE("/user/:id", u.delete)
}
