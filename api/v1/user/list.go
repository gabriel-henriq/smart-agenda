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
		Offset:  (req.CurrentPage - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	users, err := u.db.ListUsers(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(users) == 0 {
		rsp := models.Paginate(models.ListUserResponse{Users: []models.UserResponse{}}, &req, 0)
		ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, rsp))
		return
	}

	parsedUsers := models.UsersToJSONList(users)
	rsp := models.Paginate(parsedUsers.Users, &req, users[0].TotalItems)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, rsp))
}
