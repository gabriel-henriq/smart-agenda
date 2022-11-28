package aulas

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type getAulaRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (r Aula) getAulaByID(ctx *gin.Context) {
	var req getAulaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	aula, err := r.db.GetAulaByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := r.toJSONAula(aula)

	ctx.JSON(http.StatusOK, rsp)
}
