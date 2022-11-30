package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
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
	Professors                []ProfessorResponse `json:"professors"`
	models.PaginationResponse `json:"pagination"`
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
		PaginationResponse: models.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLProfessors[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
