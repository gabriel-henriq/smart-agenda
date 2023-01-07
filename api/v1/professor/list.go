package professor

import (
	"net/http"

	"github.com/gabriel-henriq/smart-agenda/models"

	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func (p Professor) list(ctx *gin.Context) {
	var req models.PaginationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.ListProfessorsParams{
		Limit:   req.PageSize,
		Offset:  (req.CurrentPage - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	profs, err := p.db.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(profs) == 0 {
		rsp := models.Paginate(models.ListProfessorsResponse{Professors: []models.ProfessorResponse{}}, &req, 0)
		ctx.JSON(http.StatusOK, models.ResponseData("200", "", true, rsp))
		return
	}

	rsp := models.ProferrosToJSONList(profs)

	ctx.JSON(http.StatusOK, models.ResponseData("200", "Professores listados com sucesso", true, rsp))
}
