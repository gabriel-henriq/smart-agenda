package api

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type listProfessorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProfessor(ctx *gin.Context) {
	var req listProfessorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.ListProfessorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	profs, err := server.store.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profs)
}
