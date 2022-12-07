package professor

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func (p Professor) update(ctx *gin.Context) {
	var req models.UpdateProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.UpdateProfessorByIDParams{
		ID:         req.ID,
		Name:       sql.NullString{String: req.Name, Valid: req.Name != ""},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: req.LabelColor != ""},
	}

	prof, err := p.db.UpdateProfessorByID(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.ProfessorToJSON(prof)

	ctx.JSON(http.StatusOK, rsp)
}
