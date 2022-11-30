package aula

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (a Aula) updateAula(ctx *gin.Context) {
	var req UpdateAulaRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := sqlc.UpdateAulaByIDParams{
		ID:          req.ID,
		TabletID:    req.TabletID,
		ProfessorID: req.ProfessorID,
		RoomID:      req.RoomID,
		StudentName: req.StudentName,
		Observation: req.Observation,
		MeetStart:   req.MeetStart,
		MeetEnd:     req.MeetEnd,
	}

	aula, err := a.db.UpdateAulaByID(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := ToJSONAula(aula)

	ctx.JSON(http.StatusOK, rsp)
}
