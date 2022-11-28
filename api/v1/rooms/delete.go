package room

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type deleteRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (r Room) deleteRoom(ctx *gin.Context) {
	var req deleteRoomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	err := r.db.DeleteRoomByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
