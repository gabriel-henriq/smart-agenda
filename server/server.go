package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/api/v1/aula"
	"github.com/gabriel-henriq/smart-agenda/api/v1/paseto"
	"github.com/gabriel-henriq/smart-agenda/api/v1/professor"
	"github.com/gabriel-henriq/smart-agenda/api/v1/room"
	"github.com/gabriel-henriq/smart-agenda/api/v1/tablet"
	"github.com/gabriel-henriq/smart-agenda/api/v1/user"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) *Server {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	server.createRoutesV1(router)

	server.router = router
}

func (server *Server) createRoutesV1(router *gin.Engine) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	v1 := router.Group("/v1")

	professorRoutes := professor.NewProfessor(server.store, server.config)
	roomRoutes := room.NewRoom(server.store, server.config)
	tabletRoutes := tablet.NewTablet(server.store, server.config)
	aulaRoutes := aula.NewAula(server.store, server.config)
	userRoutes := user.NewUser(server.store, server.config)
	tokenRoutes := paseto.NewToken(server.store)

	professorRoutes.SetupProfessorRoute(v1)
	roomRoutes.SetupRoomRoute(v1)
	tabletRoutes.SetupTabletRoute(v1)
	aulaRoutes.SetupAulaRoute(v1)
	userRoutes.SetupUserRoute(v1)
	tokenRoutes.SetupTokenRoute(v1)

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
