package room

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func (r Room) list(ctx *gin.Context) {
	var req models.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := sqlc.ListRoomsParams{
		Limit:   req.PageSize,
		Offset:  (req.CurrentPage - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	rooms, err := r.db.ListRooms(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(rooms) == 0 {
		rsp := models.Paginate(models.ListRoomResponse{Rooms: []models.RoomResponse{}}, &req, 0)
		ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, rsp))
		return
	}

	rsp := models.RoomsToJSONList(rooms)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "Salas listadas com sucesso", true, rsp))
}
