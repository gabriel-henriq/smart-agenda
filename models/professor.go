package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

type CreateProfessorRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type DeleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type UpdateProfessorRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type ProfessorResponse struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type ListProfessorsResponse struct {
	Professors         []ProfessorResponse `json:"professors"`
	PaginationResponse `json:"pagination"`
}

func ProfessorToJSON(sqlProfessor sqlc.Professor) ProfessorResponse {
	return ProfessorResponse{
		ID:         sqlProfessor.ID,
		Name:       sqlProfessor.Name,
		LabelColor: sqlProfessor.LabelColor,
		CreatedAt:  sqlProfessor.CreatedAt.Unix(),
		UpdatedAt:  sqlProfessor.UpdatedAt.Unix(),
	}
}

func ProferrosToJSONList(SQLProfessors []sqlc.ListProfessorsRow) ListProfessorsResponse {
	var profs []ProfessorResponse

	for _, professor := range SQLProfessors {
		profs = append(profs, ProfessorResponse{
			ID:         professor.ID,
			Name:       professor.Name,
			LabelColor: professor.LabelColor,
			CreatedAt:  professor.CreatedAt.Unix(),
			UpdatedAt:  professor.UpdatedAt.Unix(),
		})
	}

	return ListProfessorsResponse{Professors: profs}
}
