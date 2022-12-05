package models

type PaginationResponse struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"skip"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int32 `json:"totalPages"`
}

type PaginationRequest struct {
	PageSize int32  `form:"pageSize"`
	PageID   int32  `form:"pageId"`
	OrderBy  string `form:"orderBy"`
	Reverse  bool   `form:"reverse"`
}
