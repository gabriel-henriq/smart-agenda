package professor

import (
	"github.com/gabriel-henriq/smart-agenda/models"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (p Professor) list(ctx *gin.Context) {
	var req models.PaginationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.ListProfessorsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	profs, err := p.db.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(profs) == 0 {
		ctx.JSON(http.StatusOK, models.ListProfessorsResponse{
			Professors: []models.ResponseProfessor{},
			PaginationResponse: models.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := models.ProferrosToJSONList(profs, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
