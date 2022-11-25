package professor

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

//type professorResponse struct {
//	Name       string `json:"name" binding:"omitempty"`
//	LabelColor string `json:"labelColor" binding:"omitempty"`
//}

func (p Professor) getProfessor(ctx *gin.Context) {
	var req getProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}

	prof, err := p.GetProfessorByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, api.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	rsp := toJSONProfessor(prof)

	ctx.JSON(http.StatusOK, rsp)
}