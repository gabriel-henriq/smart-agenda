package room

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (r Room) update(ctx *gin.Context) {
	var req updateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.UpdateRoomByIDParams{
		ID:         req.ID,
		Name:       sql.NullString{String: req.Name, Valid: req.Name != ""},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: req.LabelColor != ""},
	}

	room, err := r.db.UpdateRoomByID(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := toJSON(room)

	ctx.JSON(http.StatusOK, rsp)
}
