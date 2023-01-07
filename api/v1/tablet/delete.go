package tablet

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
)

func (t Tablet) delete(ctx *gin.Context) {
	var req models.DeleteTabletRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := t.db.DeleteTabletByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseData("500", err.Error(), false, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, nil))
}
