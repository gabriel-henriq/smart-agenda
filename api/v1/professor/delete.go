package professor

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"
)

func (p Professor) delete(ctx *gin.Context) {
	var req models.DeleteProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := p.db.DeleteProfessorByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseData("500", err.Error(), false, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, nil))
}
