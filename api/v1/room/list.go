package room

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (r Room) list(ctx *gin.Context) {
	var req v1.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(rooms) == 0 {
		ctx.JSON(http.StatusOK, listResponse{
			Rooms: []response{},
			PaginationResponse: v1.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := toJSONList(rooms, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
