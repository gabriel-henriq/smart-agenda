package aula

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

func (a Aula) getAulaByID(ctx *gin.Context) {
	var req models.GetAulaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	aula, err := a.db.GetAulaByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := models.ToJSONAula(aula)

	ctx.JSON(http.StatusOK, rsp)
}
