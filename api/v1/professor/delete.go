package professor

import (
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
