package tablet

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (t Tablet) update(ctx *gin.Context) {
	var req updateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	arg := sqlc.UpdateTabletByIDParams{
		ID:         req.ID,
		Name:       sql.NullString{String: req.Name, Valid: req.Name != ""},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: req.LabelColor != ""},
	}

	prof, err := t.db.UpdateTabletByID(ctx, arg)
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

	rsp := toJSON(prof)

	ctx.JSON(http.StatusOK, rsp)
}
