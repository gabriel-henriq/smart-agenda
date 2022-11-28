package professors

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type deleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (p Professor) deleteProfessor(ctx *gin.Context) {
	var req deleteProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	err := p.db.DeleteProfessorByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
