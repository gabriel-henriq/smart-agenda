package tablet

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

func (t Tablet) getTabletByID(ctx *gin.Context) {
	var req models.GetTabletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	room, err := t.db.GetTabletByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := models.ToJSONTablet(room)

	ctx.JSON(http.StatusOK, rsp)
}
