package user

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (u User) update(ctx *gin.Context) {
	var req models.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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
				ctx.JSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.UserToJSON(user)

	ctx.JSON(http.StatusOK, rsp)
}
