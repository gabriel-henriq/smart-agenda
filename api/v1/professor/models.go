package professor

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
	"math"
)

func (p Professor) toJSONProfessor(sqlProfessor sqlc.Professor) models.Professor {
	return models.Professor{
		Name:       sqlProfessor.Name.String,
		LabelColor: sqlProfessor.LabelColor.String,
	}
}

func (p Professor) toJSONProfessorList(SQLProfessors []sqlc.ListProfessorsRow, pageID, pageSize int32) models.ProfessorList {
	var profs []models.Professor

	for _, sqlP := range SQLProfessors {
		profs = append(profs, models.Professor{
			Name:       sqlP.Name.String,
			LabelColor: sqlP.LabelColor.String,
		})
	}

	totalPages := int32(math.Ceil(math.Round(float64(pageID / pageSize))))

	return models.ProfessorList{
		Professors: profs,
		Pagination: models.Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalPages: totalPages,
		},
	}
}
