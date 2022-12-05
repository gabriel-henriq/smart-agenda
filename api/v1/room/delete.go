package room

import (
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Room) delete(ctx *gin.Context) {
	var req models.DeleteRoomRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := r.db.DeleteRoomByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
