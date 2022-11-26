package server

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1/professor"
	"github.com/gabriel-henriq/smart-agenda/api/v1/rooms"
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

	professorRoutes := professor.NewProfessor(server.store)
	roomRoutes := room.NewRoom(server.store)

	professorRoutes.SetupProfessorRoute(router)
	roomRoutes.SetupRoomRoute(router)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
