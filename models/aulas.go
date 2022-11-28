package models

import (
	"time"
)

type Aula struct {
	ID          int32     `json:"ID"`
	TabletID    int32     `json:"tabletID"`
	ProfessorID int32     `json:"professorID"`
	RoomID      int32     `json:"roomID"`
	StudentName string    `json:"studentName"`
	Observation string    `json:"observation"`
	MeetStart   time.Time `json:"meetStart"`
	MeetEnd     time.Time `json:"meetEnd"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AulaDetails struct {
	ID            int32     `json:"id"`
	StudentName   string    `json:"studentName"`
	ProfessorName string    `json:"professorName"`
	TabletName    string    `json:"tabletName"`
	RoomName      string    `json:"roomName"`
	MeetStart     time.Time `json:"meetStart"`
	MeetEnd       time.Time `json:"meetEnd"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
