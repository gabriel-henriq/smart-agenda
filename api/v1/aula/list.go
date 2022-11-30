package aula

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Aula) listAula(ctx *gin.Context) {
	var req ListAulaRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	args := sqlc.ListAulasByTimeRangeParams{
		MeetStart: req.MeetStart,
		MeetEnd:   req.MeetEnd,
	}

	aulas, err := a.db.ListAulasByTimeRange(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(aulas) == 0 {
		ctx.JSON(http.StatusOK, ListAulaResponse{
			Aulas: []AulaResponse{},
		})
		return
	}

	rsp := ToJSONAulasList(aulas)

	ctx.JSON(http.StatusOK, rsp)
}
