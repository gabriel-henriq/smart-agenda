package aula

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a Aula) delete(ctx *gin.Context) {
	var req deleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}
	err := a.db.DeleteAulaByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
