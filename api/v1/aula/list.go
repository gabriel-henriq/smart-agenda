package aula

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Aula) list(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	args := sqlc.ListAulasByTimeRangeParams{
		MeetStart: req.MeetStart,
		MeetEnd:   req.MeetEnd,
	}

	aulas, err := a.db.ListAulasByTimeRange(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}
	if len(aulas) == 0 {
		ctx.JSON(http.StatusOK, ListResponse{
			Aulas: []response{},
		})
		return
	}

	rsp := ToJSONList(aulas)

	ctx.JSON(http.StatusOK, rsp)
}
