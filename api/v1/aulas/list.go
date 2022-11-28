package aulas

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
)

type listAulaRequest struct {
	MeetStart time.Time `json:"meetStart" binding:"required"`
	MeetEnd   time.Time `json:"meetEnd" binding:"required"`
}

func (r Aula) listAula(ctx *gin.Context) {
	var req listAulaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	args := sqlc.ListAulasByTimeRangeParams{
		MeetStart: sql.NullTime{Time: req.MeetStart, Valid: true},
		MeetEnd:   sql.NullTime{Time: req.MeetEnd, Valid: true},
	}

	aulas, err := r.db.ListAulasByTimeRange(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(aulas) == 0 {
		ctx.JSON(http.StatusOK, []models.AulaDetails{})
		return
	}

	rsp := r.toJSONAulaList(aulas)

	ctx.JSON(http.StatusOK, rsp)
}
