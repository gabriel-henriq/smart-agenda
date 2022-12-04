package professor

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p Professor) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	prof, err := p.db.GetProfessorByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, v1.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := toJSON(prof)

	ctx.JSON(http.StatusOK, rsp)
}
