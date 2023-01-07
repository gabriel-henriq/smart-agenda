package aula

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
)

func (a Aula) list(ctx *gin.Context) {
	var req models.ListAulasRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := sqlc.ListAulasByTimeRangeParams{
		MeetStart: req.MeetStart,
		MeetEnd:   req.MeetEnd,
	}

	aulas, err := a.db.ListAulasByTimeRange(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(aulas) == 0 {
		ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, models.ListAulasResponse{Aulas: []models.AulaResponse{}}))
		return
	}

	rsp := models.AulasToJSONList(aulas)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "Aulas listadas com sucesso", true, rsp))
}
