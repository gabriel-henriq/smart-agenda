package aula

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

func (a Aula) create(ctx *gin.Context) {
	var req createRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateAulaParams{
		TabletID:    req.TabletID,
		ProfessorID: req.ProfessorID,
		RoomID:      req.RoomID,
		StudentName: req.StudentName,
		Observation: req.Observation,
		MeetStart:   req.MeetStart,
		MeetEnd:     req.MeetEnd,
	}

	aula, err := a.db.CreateAula(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, v1.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := toJSON(aula)

	ctx.JSON(http.StatusOK, rsp)
}
