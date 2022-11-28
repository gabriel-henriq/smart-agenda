package tablets

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type listTabletRequest struct {
	PageSize int32  `form:"page_size" binding:"required,min=1,max=100"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	OrderBy  string `form:"order_by"`
	Reverse  bool   `form:"reverse"`
}

func (r Tablet) listTablet(ctx *gin.Context) {
	var req listTabletRequest
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

	rooms, err := r.db.ListTablets(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(rooms) == 0 {
		ctx.JSON(http.StatusOK, models.TabletList{
			Tablets: []models.Tablet{},
			Pagination: models.Pagination{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := r.toJSONTabletList(rooms, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
