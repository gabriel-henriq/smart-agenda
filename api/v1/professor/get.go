package professor

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (p Professor) getProfessor(ctx *gin.Context) {
	var req getProfessorRequest
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

	rsp := p.toJSONProfessor(prof)

	ctx.JSON(http.StatusOK, rsp)
}
