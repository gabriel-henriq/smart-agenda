package tablet

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

func (t Tablet) listTablet(ctx *gin.Context) {
	var req models.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	args := sqlc.ListTabletsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	rooms, err := t.db.ListTablets(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(rooms) == 0 {
		ctx.JSON(http.StatusOK, TabletList{
			Tablets: []TabletResponse{},
			PaginationResponse: models.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := ToJSONTabletList(rooms, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
