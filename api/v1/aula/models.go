package aula

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"time"
)

type createRequest struct {
	TabletID    int32     `json:"tabletId" binding:"numeric"`
	ProfessorID int32     `json:"professorId" binding:"required,numeric"`
	RoomID      int32     `json:"roomId" binding:"required,numeric"`
	StudentName string    `json:"studentName" binding:"alpha"`
	MeetStart   time.Time `json:"meetStart" binding:"required,ltefield=MeetEnd"`
	MeetEnd     time.Time `json:"meetEnd" binding:"required"`
	Observation string    `json:"observation"`
}

type updateRequest struct {
	ID          int32     `json:"id" binding:"required,numeric"`
	TabletID    int32     `json:"tabletId" binding:"numeric"`
	ProfessorID int32     `json:"professorId" binding:"numeric"`
	RoomID      int32     `json:"roomId" binding:"numeric"`
	StudentName string    `json:"studentName" binding:"alpha"`
	MeetStart   time.Time `json:"meetStart" binding:"ltefield=MeetEnd"`
	MeetEnd     time.Time `json:"meetEnd" binding:""`
	Observation string    `json:"observation"`
}

type response struct {
	ID          int32     `json:"id"`
	TabletID    int32     `json:"TabletId"`
	ProfessorID int32     `json:"ProfessorId"`
	RoomID      int32     `json:"RoomId"`
	StudentName string    `json:"studentName"`
	Observation string    `json:"observation"`
	MeetStart   time.Time `json:"meetStart"`
	MeetEnd     time.Time `json:"meetEnd"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type listResponse struct {
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

type deleteRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type GetRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listRequest struct {
	MeetStart time.Time `form:"meetStart" binding:"ltefield=MeetEnd"`
	MeetEnd   time.Time `form:"meetEnd"`
}

type ListResponse struct {
	Aulas []response `json:"Aulas"`
}

func ToJSON(sqlAula sqlc.Aula) response {
	return response{
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

func ToJSONList(SQLAulas []sqlc.ListAulasByTimeRangeRow) []listResponse {
	var aulas []listResponse

	for _, aula := range SQLAulas {
		aulas = append(aulas, listResponse{
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
