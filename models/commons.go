package models

type Pagination struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"skip"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int32 `json:"totalPages"`
}
