package aulas

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAulaRequest struct {
	TabletID    int32     `json:"tabletID" binding:"numeric"`
	ProfessorID int32     `json:"professorID" binding:"required,numeric"`
	RoomID      int32     `json:"roomID" binding:"required,numeric"`
	StudentName string    `json:"studentName" binding:"alpha"`
	Observation string    `json:"observation" binding:"len=255"`
	MeetStart   time.Time `json:"meetStart" binding:"required,ltefield=MeetEnd" time_format:"2022-11-28T01:28:09.995389Z"`
	MeetEnd     time.Time `json:"meetEnd" binding:"required" time_format:"2022-11-28T01:28:09.995389Z"`
}

func (r Aula) createAula(ctx *gin.Context) {
	var req models.Aula

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateAulaParams{
		TabletID:    sql.NullInt32{Int32: req.TabletID, Valid: true},
		ProfessorID: sql.NullInt32{Int32: req.ProfessorID, Valid: true},
		RoomID:      sql.NullInt32{Int32: req.RoomID, Valid: true},
		StudentName: sql.NullString{String: req.StudentName, Valid: true},
		Observation: sql.NullString{String: req.Observation, Valid: true},
		MeetStart:   sql.NullTime{Time: req.MeetStart, Valid: true},
		MeetEnd:     sql.NullTime{Time: req.MeetEnd, Valid: true},
	}

	aula, err := r.db.CreateAula(ctx, arg)
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

	rsp := r.toJSONAula(aula)

	ctx.JSON(http.StatusOK, rsp)
}
