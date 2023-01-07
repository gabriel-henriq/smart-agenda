package aula

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
)

func (a Aula) getByID(ctx *gin.Context) {
	var req models.GetAulaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	aula, err := a.db.GetAulaByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.AulaToJSON(aula)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "Aula listada com sucesso", true, rsp))
}
