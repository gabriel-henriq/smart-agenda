package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p Professor) delete(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}
	err := p.db.DeleteProfessorByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
