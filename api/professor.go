package api

import (
	"database/sql"
	db "github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createProfessorRequest struct {
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

func (server *Server) createProfessor(ctx *gin.Context) {
	var req createProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	profParams := db.CreateProfessorParams{
		Name:       sql.NullString{String: req.Name, Valid: true},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: true},
	}

	prof, _ := server.store.CreateProfessor(ctx, profParams)

	ctx.JSON(http.StatusOK, prof)
}
