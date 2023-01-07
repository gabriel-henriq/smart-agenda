package tablet

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
)

func (t Tablet) getByID(ctx *gin.Context) {
	var req models.GetTabletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	room, err := t.db.GetTabletByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.TabletToJSON(room)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, rsp))
}
