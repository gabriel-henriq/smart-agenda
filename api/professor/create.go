package professor

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createProfessorRequest struct {
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

func toJSONProfessor(professor sqlc.Professor) professorResponse {
	return professorResponse{
		Name:       professor.Name.String,
		LabelColor: professor.LabelColor.String,
	}
}

func (p Professor) createProfessor(ctx *gin.Context) {
	var req createProfessorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateProfessorParams{
		Name:       sql.NullString{String: req.Name, Valid: true},
		LabelColor: sql.NullString{String: req.LabelColor, Valid: true},
	}

	prof, err := p.CreateProfessor(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, api.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(err))
		return
	}

	rsp := toJSONProfessor(prof)

	ctx.JSON(http.StatusOK, rsp)
}
