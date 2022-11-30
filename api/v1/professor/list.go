package professor

import (
	"github.com/gabriel-henriq/smart-agenda/utils"
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gin-gonic/gin"
)

func (p Professor) listProfessor(ctx *gin.Context) {
	var req models.ListProfessorRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if len(profs) == 0 {
		ctx.JSON(http.StatusOK, models.ProfessorListResponse{
			Professors: []models.ProfessorResponse{},
			Pagination: models.Pagination{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := models.ToJSONProfessorList(profs, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
