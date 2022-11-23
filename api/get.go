package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProfessor(ctx *gin.Context) {
	var req getProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	prof, err := server.store.GetProfessorByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newProfessorResponse(prof)

	ctx.JSON(http.StatusOK, rsp)
}
