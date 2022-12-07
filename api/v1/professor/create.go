package professor

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func (p Professor) create(ctx *gin.Context) {
	var req models.CreateProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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
