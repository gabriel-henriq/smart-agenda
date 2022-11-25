package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (p Professor) deleteProfessor(ctx *gin.Context) {
	var req deleteProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}
	err := p.DeleteProfessorByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
