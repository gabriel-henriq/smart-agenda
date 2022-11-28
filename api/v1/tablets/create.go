package tablets

import (
	"database/sql"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (r Tablet) createTablet(ctx *gin.Context) {
	var req models.Tablet

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := sqlc.Tablet{
		Name: sql.NullString{String: req.Name, Valid: true},
	}

	tablet, err := r.db.CreateTablet(ctx, arg.Name)
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

	rsp := r.toJSONTablet(tablet)

	ctx.JSON(http.StatusOK, rsp)
}
