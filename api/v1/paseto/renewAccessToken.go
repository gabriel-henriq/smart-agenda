package paseto

import (
	"database/sql"
	"fmt"
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (t Token) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, v1.ErrorResponse(err))
		return
	}

	refreshPayload, err := t.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, v1.ErrorResponse(err))
		return
	}

	session, err := t.db.GetSessionByID(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, v1.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, v1.ErrorResponse(err))
		return
	}

	if session.Email != refreshPayload.Email {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, v1.ErrorResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, v1.ErrorResponse(err))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, v1.ErrorResponse(err))
		return
	}

	accessToken, accessPayload, err := t.tokenMaker.CreateToken(
		refreshPayload.Email,
		t.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, v1.ErrorResponse(err))
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
