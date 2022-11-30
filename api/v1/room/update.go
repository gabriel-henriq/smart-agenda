package room

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (r Room) updateRoom(ctx *gin.Context) {
	var req UpdateRoomRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := sqlc.UpdateRoomByIDParams{
		ID:         req.ID,
		Name:       req.Name,
		LabelColor: req.LabelColor,
	}

	room, err := r.db.UpdateRoomByID(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := ToJSONRoom(room)

	ctx.JSON(http.StatusOK, rsp)
}
