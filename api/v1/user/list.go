package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
)

func (u User) list(ctx *gin.Context) {
	var req models.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := sqlc.ListUsersParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	users, err := u.db.ListUsers(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusOK, models.ListUserResponse{
			Users: []models.UserResponse{},
			PaginationResponse: models.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := models.UsersToJSONList(users, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
