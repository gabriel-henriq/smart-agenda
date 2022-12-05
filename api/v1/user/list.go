package user

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u User) list(ctx *gin.Context) {
	var req v1.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusOK, listResponse{
			Users: []response{},
			PaginationResponse: v1.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := toJSONList(users, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
