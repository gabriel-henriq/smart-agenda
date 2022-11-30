package aula

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"time"
)

type CreateAulaRequest struct {
	TabletID    int32     `json:"tabletID" binding:"numeric"`
	ProfessorID int32     `json:"professorID" binding:"required,numeric"`
	RoomID      int32     `json:"roomID" binding:"required,numeric"`
	StudentName string    `json:"studentName" binding:"alpha"`
	MeetStart   time.Time `json:"meetStart" binding:"required,ltefield=MeetEnd"`
	MeetEnd     time.Time `json:"meetEnd" binding:"required"`
	Observation string    `json:"observation"`
}

type UpdateAulaRequest struct {
	ID          int32     `json:"id" binding:"required,numeric"`
	TabletID    int32     `json:"tabletID" binding:"numeric"`
	ProfessorID int32     `json:"professorID" binding:"numeric"`
	RoomID      int32     `json:"roomID" binding:"numeric"`
	StudentName string    `json:"studentName" binding:"alpha"`
	MeetStart   time.Time `json:"meetStart" binding:"ltefield=MeetEnd"`
	MeetEnd     time.Time `json:"meetEnd" binding:""`
	Observation string    `json:"observation"`
}

type AulaResponse struct {
	ID          int32     `json:"id"`
	TabletID    int32     `json:"TabletID"`
	ProfessorID int32     `json:"ProfessorID"`
	RoomID      int32     `json:"RoomID"`
	StudentName string    `json:"studentName"`
	Observation string    `json:"observation"`
	MeetStart   time.Time `json:"meetStart"`
	MeetEnd     time.Time `json:"meetEnd"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AulasResponse struct {
	ID            int32     `json:"id"`
	StudentName   string    `json:"studentName"`
	MeetStart     time.Time `json:"meetStart"`
	MeetEnd       time.Time `json:"meetEnd"`
	ProfessorName string    `json:"professorName"`
	TabletName    string    `json:"tabletName"`
	RoomName      string    `json:"observation"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type DeleteAulaRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type GetAulaRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type ListAulaRequest struct {
	MeetStart time.Time `form:"meet_start" binding:"ltefield=MeetEnd"`
	MeetEnd   time.Time `form:"meet_end"`
}

type ListAulaResponse struct {
	Aulas []AulaResponse `json:"Aulas"`
}

func ToJSONAula(sqlAula sqlc.Aula) AulaResponse {
	return AulaResponse{
		ID:          sqlAula.ID,
		TabletID:    sqlAula.TabletID,
		ProfessorID: sqlAula.ProfessorID,
		RoomID:      sqlAula.RoomID,
		StudentName: sqlAula.StudentName,
		Observation: sqlAula.Observation,
		MeetStart:   sqlAula.MeetStart,
		MeetEnd:     sqlAula.MeetEnd,
		CreatedAt:   sqlAula.CreatedAt,
		UpdatedAt:   sqlAula.UpdatedAt,
	}
}

func ToJSONAulasList(SQLAulas []sqlc.ListAulasByTimeRangeRow) []AulasResponse {
	var aulas []AulasResponse

	for _, aula := range SQLAulas {
		aulas = append(aulas, AulasResponse{
			ID:            aula.ID,
			StudentName:   aula.StudentName,
			MeetStart:     aula.MeetStart,
			MeetEnd:       aula.MeetEnd,
			ProfessorName: aula.ProfessorName,
			TabletName:    aula.TabletName,
			RoomName:      aula.RoomName,
			CreatedAt:     aula.CreatedAt,
			UpdatedAt:     aula.UpdatedAt,
		})
	}

	return aulas
}
