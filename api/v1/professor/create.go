package professor

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (p Professor) createProfessor(ctx *gin.Context) {
	var req CreateProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateProfessorParams{
		Name:       req.Name,
		LabelColor: req.LabelColor,
	}

	prof, err := p.db.CreateProfessor(ctx, arg)
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

	rsp := ToJSONProfessor(prof)

	ctx.JSON(http.StatusOK, rsp)
}
