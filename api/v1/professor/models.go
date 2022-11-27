package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"math"
)

func (p Professor) toJSONProfessor(sqlProfessor sqlc.Professor) models.Professor {
	return models.Professor{
		ID:         sqlProfessor.ID,
		Name:       sqlProfessor.Name.String,
		LabelColor: sqlProfessor.LabelColor.String,
		CreatedAt:  sqlProfessor.CreatedAt.Time,
		UpdatedAt:  sqlProfessor.UpdatedAt.Time,
	}
}

func (p Professor) toJSONProfessorList(SQLProfessors []sqlc.ListProfessorsRow, pageID, pageSize int32) models.ProfessorList {
	var profs []models.Professor

	for _, professor := range SQLProfessors {
		profs = append(profs, models.Professor{
			ID:         professor.ID,
			Name:       professor.Name.String,
			LabelColor: professor.LabelColor.String,
			CreatedAt:  professor.CreatedAt.Time,
			UpdatedAt:  professor.UpdatedAt.Time,
		})
	}

	totalPages := int32(math.Ceil(float64(SQLProfessors[0].TotalItems) / float64(pageSize)))

	return models.ProfessorList{
		Professors: profs,
		Pagination: models.Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLProfessors[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
