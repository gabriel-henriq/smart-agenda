package api

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createProfessorRequest struct {
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type professorResponse struct {
	Name       string `json:"name" binding:"omitempty"`
	LabelColor string `json:"labelColor" binding:"omitempty"`
}

func newProfessorResponse(professor sqlc.Professor) professorResponse {
	return professorResponse{
		Name:       professor.Name.String,
		LabelColor: professor.LabelColor.String,
	}
}

func (server *Server) createProfessor(ctx *gin.Context) {
	var req createProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.CreateProfessorParams{
		Name:       sql.NullString{String: req.Name, Valid: true},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: true},
	}

	prof, err := server.store.CreateProfessor(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newProfessorResponse(prof)

	ctx.JSON(http.StatusOK, rsp)
}
