package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

func toJSONProfessor(professor sqlc.Professor) professorResponse {
	return professorResponse{
		Name:       professor.Name.String,
		LabelColor: professor.LabelColor.String,
	}
}

func toJSONProfessorArray(professors []sqlc.ListProfessorsRow) []professorResponse {
	resps := make([]professorResponse, 0)
	for _, professor := range professors {
		resps = append(resps, professorResponse{
			Name:       professor.Name.String,
			LabelColor: professor.LabelColor.String,
		})
	}
	resps = append(resps)
	return resps
}

type getProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listProfessorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

type deleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type createProfessorRequest struct {
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type professorResponse struct {
	Name       string `json:"name" binding:"omitempty"`
	LabelColor string `json:"labelColor" binding:"omitempty"`
	pagination
}

type pagination struct {
	TotalItems int64 `json:"totalItems" binding:"omitempty"`
	Limit      int64 `json:"limit" binding:"omitempty"`
	Skip       int64 `json:"skip" binding:"omitempty"`
}
