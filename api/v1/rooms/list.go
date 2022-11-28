package room

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type listRoomRequest struct {
	PageSize int32  `form:"page_size" binding:"required,min=1,max=100"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	OrderBy  string `form:"order_by"`
	Reverse  bool   `form:"reverse"`
}

func (r Room) listRoom(ctx *gin.Context) {
	var req listRoomRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	args := sqlc.ListRoomsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	rooms, err := r.db.ListRooms(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(rooms) == 0 {
		ctx.JSON(http.StatusOK, models.RoomList{
			Rooms: []models.Room{},
			Pagination: models.Pagination{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := r.toJSONRoomList(rooms, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
