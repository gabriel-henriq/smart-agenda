package professor

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type createRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type deleteRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type updateRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type response struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

type listResponse struct {
	Professors            []response `json:"professors"`
	v1.PaginationResponse `json:"pagination"`
}

func toJSON(sqlProfessor sqlc.Professor) response {
	return response{
		ID:         sqlProfessor.ID,
		Name:       sqlProfessor.Name,
		LabelColor: sqlProfessor.LabelColor,
		CreatedAt:  sqlProfessor.CreatedAt.Unix(),
		UpdatedAt:  sqlProfessor.UpdatedAt.Unix(),
	}
}

func toJSONList(SQLProfessors []sqlc.ListProfessorsRow, pageID, pageSize int32) listResponse {
	var profs []response

	for _, professor := range SQLProfessors {
		profs = append(profs, response{
			ID:         professor.ID,
			Name:       professor.Name,
			LabelColor: professor.LabelColor,
			CreatedAt:  professor.CreatedAt.Unix(),
			UpdatedAt:  professor.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLProfessors[0].TotalItems) / float64(pageSize)))

	return listResponse{
		Professors: profs,
		PaginationResponse: v1.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLProfessors[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
