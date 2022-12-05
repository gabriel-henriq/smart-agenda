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

type deleteRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listRequest struct {
	MeetStart time.Time `form:"meetStart" binding:"ltefield=MeetEnd"`
	MeetEnd   time.Time `form:"meetEnd"`
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
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
}

type listResponse struct {
	Aulas []response `json:"aulas"`
}

func toJSON(sqlAula sqlc.Aula) response {
	return response{
		ID:          sqlAula.ID,
		TabletID:    sqlAula.TabletID,
		ProfessorID: sqlAula.ProfessorID,
		RoomID:      sqlAula.RoomID,
		StudentName: sqlAula.StudentName,
		Observation: sqlAula.Observation,
		MeetStart:   sqlAula.MeetStart,
		MeetEnd:     sqlAula.MeetEnd,
		CreatedAt:   sqlAula.CreatedAt.Unix(),
		UpdatedAt:   sqlAula.UpdatedAt.Unix(),
	}
}

func toJSONList(SQLAulas []sqlc.ListAulasByTimeRangeRow) listResponse {
	var aulas []response

	for _, aula := range SQLAulas {
		aulas = append(aulas, response{
			ID:          aula.ID,
			StudentName: aula.StudentName,
			MeetStart:   aula.MeetStart,
			MeetEnd:     aula.MeetEnd,
			ProfessorID: aula.ProfessorID,
			Observation: aula.Observation,
			TabletID:    aula.TabletID,
			RoomID:      aula.RoomID,
			CreatedAt:   aula.CreatedAt.Unix(),
			UpdatedAt:   aula.UpdatedAt.Unix(),
		})
	}

	return listResponse{Aulas: aulas}
}
