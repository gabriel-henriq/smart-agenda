package tablets

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type deleteTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (r Tablet) deleteTablet(ctx *gin.Context) {
	var req deleteTabletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	err := r.db.DeleteTabletByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
