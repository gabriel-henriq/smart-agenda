package models

import (
	"time"
)

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
