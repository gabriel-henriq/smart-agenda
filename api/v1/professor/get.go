package professor

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/models"
)

func (p Professor) get(ctx *gin.Context) {
	var req models.GetProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	prof, err := p.db.GetProfessorByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.ProfessorToJSON(prof)

	ctx.JSON(http.StatusOK, rsp)
}
