package server

import (
	"github.com/gabriel-henriq/smart-agenda/api/professor"
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

	professorRoutes.SetupProfessorRoute(router)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
