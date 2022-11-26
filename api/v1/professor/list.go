package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type listProfessorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=100"`
}

func (p Professor) listProfessor(ctx *gin.Context) {
	var req listProfessorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}

	arg := sqlc.ListProfessorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	profs, err := p.db.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}
	if len(profs) == 0 {
		ctx.JSON(http.StatusOK, models.ProfessorList{
			Professors: []models.Professor{},
			Pagination: models.Pagination{
				TotalItems: 0,
				Limit:      req.PageID,
				Offset:     req.PageSize,
				TotalPages: 0,
			},
		})
		return
	}

	rsp := p.toJSONProfessorList(profs, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
