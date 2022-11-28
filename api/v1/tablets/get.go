package tablets

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type getTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (r Tablet) getTabletByID(ctx *gin.Context) {
	var req getTabletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	room, err := r.db.GetTabletByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := r.toJSONTablet(room)

	ctx.JSON(http.StatusOK, rsp)
}
