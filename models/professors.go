package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type CreateProfessorRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
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
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type ProfessorListResponse struct {
	Professors []ProfessorResponse `json:"professors"`
	Pagination `json:"pagination"`
}

func ToJSONProfessor(sqlProfessor sqlc.Professor) ProfessorResponse {
	return ProfessorResponse{
		ID:         sqlProfessor.ID,
		Name:       sqlProfessor.Name,
		LabelColor: sqlProfessor.LabelColor,
		CreatedAt:  sqlProfessor.CreatedAt.String(),
		UpdatedAt:  sqlProfessor.UpdatedAt.String(),
	}
}

type DeleteProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type ListProfessorRequest struct {
	PageSize int32  `form:"pageSize"`
	PageID   int32  `form:"pageID"`
	OrderBy  string `form:"orderBy"`
	Reverse  bool   `form:"reverse"`
}

func ToJSONProfessorList(SQLProfessors []sqlc.ListProfessorsRow, pageID, pageSize int32) ProfessorListResponse {
	var profs []ProfessorResponse

	for _, professor := range SQLProfessors {
		profs = append(profs, ProfessorResponse{
			ID:         professor.ID,
			Name:       professor.Name,
			LabelColor: professor.LabelColor,
			CreatedAt:  professor.CreatedAt.String(),
			UpdatedAt:  professor.UpdatedAt.String(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLProfessors[0].TotalItems) / float64(pageSize)))

	return ProfessorListResponse{
		Professors: profs,
		Pagination: Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLProfessors[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
