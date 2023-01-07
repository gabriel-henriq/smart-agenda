package tablet

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func (t Tablet) list(ctx *gin.Context) {
	var req models.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := sqlc.ListTabletsParams{
		Limit:   req.PageSize,
		Offset:  (req.CurrentPage - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	rooms, err := t.db.ListTablets(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(rooms) == 0 {
		ctx.JSON(http.StatusOK, models.ListTabletsResponse{
			Tablets: []models.ResponseTablet{},
			PaginationResponse: models.PaginationResponse{
				PageSize:    req.PageSize,
				CurrentPage: req.CurrentPage,
			},
		})
		return
	}

	rsp := models.TabletsToJSONList(rooms, req.CurrentPage, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
