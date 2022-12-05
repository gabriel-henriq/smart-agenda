package user

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (u User) update(ctx *gin.Context) {
	var req updateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	arg := sqlc.UpdateUserByIDParams{
		ID:       req.ID,
		Name:     sql.NullString{String: req.Name, Valid: req.Name != ""},
		Email:    sql.NullString{String: req.Email, Valid: req.Email != ""},
		Password: sql.NullString{String: req.Password, Valid: req.Password != ""},
	}

	user, err := u.db.UpdateUserByID(ctx, arg)
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

	rsp := toJSON(user)

	ctx.JSON(http.StatusOK, rsp)
}
