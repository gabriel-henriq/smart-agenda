package user

import (
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u User) delete(ctx *gin.Context) {
	var req models.DeleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := u.db.DeleteUserByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
