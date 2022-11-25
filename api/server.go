package api

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/api/professor"
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

	// Professor Routes
	professorRoutes := professor.NewProfessor(server.store)

	professorRoutes.SetupProfessorRoute(router)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
