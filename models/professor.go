package models

import "time"

type Pagination struct {
	Limit      int32 `json:"limit" binding:"omitempty"`
	Offset     int32 `json:"skip" binding:"omitempty"`
	Items      int64 `json:"totalItems" binding:"omitempty"`
	TotalPages int32 `json:"totalPages" binding:"omitempty"`
}

type Professor struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name" binding:"omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	LabelColor string    `json:"labelColor" binding:"omitempty"`
}

type ProfessorList struct {
	Professors []Professor `json:"professors"`
	Pagination `json:"pagination"`
}
