package room

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (r Room) update(ctx *gin.Context) {
	var req updateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
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
				ctx.JSON(http.StatusForbidden, v1.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := toJSON(room)

	ctx.JSON(http.StatusOK, rsp)
}
