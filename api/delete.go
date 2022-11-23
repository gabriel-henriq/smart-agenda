package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProfessor(ctx *gin.Context) {
	var req deleteProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteProfessorByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
