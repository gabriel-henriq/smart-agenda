package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type listProfessorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type pagination struct {
	TotalItems int64 `json:"totalItems" binding:"omitempty"`
	Limit      int32 `json:"limit" binding:"omitempty"`
	Offset     int32 `json:"skip" binding:"omitempty"`
}

type professorResponse struct {
	Name       string `json:"name" binding:"omitempty"`
	LabelColor string `json:"labelColor" binding:"omitempty"`
}

type listProfessorResponse struct {
	Professors []professorResponse `json:"professors" binding:"omitempty"`
	Pagination pagination          `json:"pagination" binding:"omitempty"`
}

func toJSONProfessorArray(professors []sqlc.ListProfessorsRow, limit int32, offset int32) []listProfessorResponse {
	if len(professors) == 0 {
		return nil
	}

	profs := make([]listProfessorResponse, 0)
	profs = append(profs, listProfessorResponse{Pagination: pagination{
		Limit:      limit,
		Offset:     offset,
		TotalItems: professors[0].TotalProfessors}})

	prof := make([]professorResponse, 0)
	for _, professor := range professors {
		prof = append(prof, professorResponse{
			Name:       professor.Name.String,
			LabelColor: professor.LabelColor.String,
		})
	}
	profs = append(profs, listProfessorResponse{Professors: prof})
	return profs
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

	profs, err := p.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	rsp := toJSONProfessorArray(profs, arg.Limit, arg.Offset)

	ctx.JSON(http.StatusOK, rsp)
}
