package room

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gin-gonic/gin"
)

type deleteRoomRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (r Room) deleteRoom(ctx *gin.Context) {
	var req deleteRoomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}
	err := r.db.DeleteRoomByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
