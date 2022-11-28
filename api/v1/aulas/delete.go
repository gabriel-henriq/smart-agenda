package aulas

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type deleteAulaRequest struct {
	ID int32 `json:"id" uri:"id" binding:"required,min=1"`
}

func (r Aula) deleteAula(ctx *gin.Context) {
	var req deleteAulaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	err := r.db.DeleteAulaByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
