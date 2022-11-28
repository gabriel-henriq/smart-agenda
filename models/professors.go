package models

import "time"

type Professor struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	LabelColor string    `json:"labelColor"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ProfessorList struct {
	Professors []Professor `json:"professors"`
	Pagination `json:"pagination"`
}
