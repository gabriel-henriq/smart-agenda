package room

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

func (r Room) createRooms(ctx *gin.Context) {
	var req models.Room

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}

	arg := sqlc.Room{
		Name: sql.NullString{String: req.Name, Valid: true},
	}

	room, err := r.db.CreateRoom(ctx, arg.Name)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, api.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	rsp := r.toJSONRoom(room)

	ctx.JSON(http.StatusOK, rsp)
}
