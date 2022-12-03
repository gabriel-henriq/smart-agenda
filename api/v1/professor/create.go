package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (p Professor) create(ctx *gin.Context) {
	var req createRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
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
				ctx.JSON(http.StatusForbidden, v1.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := ToJSON(prof)

	ctx.JSON(http.StatusOK, rsp)
}
