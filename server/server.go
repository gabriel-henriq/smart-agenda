package server

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/aulas"
	"github.com/gabriel-henriq/smart-agenda/api/v1/professors"
	"github.com/gabriel-henriq/smart-agenda/api/v1/rooms"
	"github.com/gabriel-henriq/smart-agenda/api/v1/tablets"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	server.createRoutesV1(router)

	server.router = router
}

// Create all V1 routes
func (server *Server) createRoutesV1(router *gin.Engine) {

	v1 := router.Group("/v1")

	professorRoutes := professors.NewProfessor(server.store)
	roomRoutes := room.NewRoom(server.store)
	tabletRoutes := tablets.NewTablet(server.store)
	aulaRoutes := aulas.NewAula(server.store)

	professorRoutes.SetupProfessorRoute(v1)
	roomRoutes.SetupRoomRoute(v1)
	tabletRoutes.SetupTabletRoute(v1)
	aulaRoutes.SetupAulaRoute(v1)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
