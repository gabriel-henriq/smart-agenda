package room

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Room) getByID(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.db.GetRoomByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := toJSON(room)

	ctx.JSON(http.StatusOK, rsp)
}
