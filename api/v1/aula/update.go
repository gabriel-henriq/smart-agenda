package aula

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (a Aula) update(ctx *gin.Context) {
	var req updateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	arg := sqlc.UpdateAulaByIDParams{
		ID:          req.ID,
		TabletID:    sql.NullInt32{Int32: req.TabletID, Valid: req.TabletID != 0},
		ProfessorID: sql.NullInt32{Int32: req.ProfessorID, Valid: req.ProfessorID != 0},
		RoomID:      sql.NullInt32{Int32: req.RoomID, Valid: req.RoomID != 0},
		StudentName: sql.NullString{String: req.StudentName, Valid: req.StudentName != ""},
		Observation: sql.NullString{String: req.Observation, Valid: req.Observation != ""},
		MeetStart:   sql.NullTime{Time: req.MeetStart, Valid: !req.MeetStart.IsZero()},
		MeetEnd:     sql.NullTime{Time: req.MeetEnd, Valid: !req.MeetEnd.IsZero()},
	}

	aula, err := a.db.UpdateAulaByID(ctx, arg)
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
