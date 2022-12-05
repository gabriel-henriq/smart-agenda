package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
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

type ResponseProfessor struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type ListProfessorsResponse struct {
	Professors         []ResponseProfessor `json:"professors"`
	PaginationResponse `json:"pagination"`
}

func ProfessorToJSON(sqlProfessor sqlc.Professor) ResponseProfessor {
	return ResponseProfessor{
		ID:         sqlProfessor.ID,
		Name:       sqlProfessor.Name,
		LabelColor: sqlProfessor.LabelColor,
		CreatedAt:  sqlProfessor.CreatedAt.Unix(),
		UpdatedAt:  sqlProfessor.UpdatedAt.Unix(),
	}
}

func ProferrosToJSONList(SQLProfessors []sqlc.ListProfessorsRow, pageID, pageSize int32) ListProfessorsResponse {
	var profs []ResponseProfessor

	for _, professor := range SQLProfessors {
		profs = append(profs, ResponseProfessor{
			ID:         professor.ID,
			Name:       professor.Name,
			LabelColor: professor.LabelColor,
			CreatedAt:  professor.CreatedAt.Unix(),
			UpdatedAt:  professor.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLProfessors[0].TotalItems) / float64(pageSize)))

	return ListProfessorsResponse{
		Professors: profs,
		PaginationResponse: PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLProfessors[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
