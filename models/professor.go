package models

type Pagination struct {
	TotalItems int64 `json:"totalItems" binding:"omitempty"`
	Limit      int32 `json:"limit" binding:"omitempty"`
	Offset     int32 `json:"skip" binding:"omitempty"`
	TotalPages int32 `json:"totalPages" binding:"omitempty"`
}

type Professor struct {
	ID         string `json:"id"`
	Name       string `json:"name" binding:"omitempty"`
	LabelColor string `json:"labelColor" binding:"omitempty"`
}

type ProfessorList struct {
	Professors []Professor `json:"professors"`
	Pagination `json:"pagination"`
}
