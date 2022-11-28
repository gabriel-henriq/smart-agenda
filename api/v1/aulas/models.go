package aulas

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
)

func (r Aula) toJSONAula(sqlAula sqlc.Aula) models.Aula {
	return models.Aula{
		ID:          sqlAula.ID,
		TabletID:    sqlAula.TabletID.Int32,
		ProfessorID: sqlAula.ProfessorID.Int32,
		RoomID:      sqlAula.RoomID.Int32,
		StudentName: sqlAula.StudentName.String,
		Observation: sqlAula.Observation.String,
		MeetStart:   sqlAula.MeetStart.Time,
		MeetEnd:     sqlAula.MeetEnd.Time,
		CreatedAt:   sqlAula.CreatedAt.Time,
		UpdatedAt:   sqlAula.UpdatedAt.Time,
	}
}

func (r Aula) toJSONAulaList(SQLAulas []sqlc.ListAulasByTimeRangeRow) []models.AulaDetails {
	var aulas []models.AulaDetails

	for _, aula := range SQLAulas {
		aulas = append(aulas, models.AulaDetails{
			ID:            aula.ID,
			StudentName:   aula.StudentName.String,
			MeetStart:     aula.MeetStart.Time,
			MeetEnd:       aula.MeetEnd.Time,
			ProfessorName: aula.ProfessorName.String,
			TabletName:    aula.TabletName.String,
			RoomName:      aula.RoomName.String,
			CreatedAt:     aula.CreatedAt.Time,
			UpdatedAt:     aula.UpdatedAt.Time,
		})
	}

	return aulas
}
