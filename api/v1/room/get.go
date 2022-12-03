package room

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Room) getByID(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	room, err := r.db.GetRoomByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, v1.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := toJSON(room)

	ctx.JSON(http.StatusOK, rsp)
}
