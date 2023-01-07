package models

import "math"

type PaginationResponse struct {
	PageSize    int32       `json:"pageSize"`
	CurrentPage int32       `json:"currentPage"`
	TotalItems  int64       `json:"totalItems"`
	TotalPages  int32       `json:"totalPages"`
	Items       interface{} `json:"items"`
}

type PaginationRequest struct {
	PageSize    int32  `form:"pageSize,default=10"`
	CurrentPage int32  `form:"currentPage,default=1"`
	OrderBy     string `form:"orderBy"`
	Reverse     bool   `form:"reverse"`
}

func Paginate(items interface{}, req *PaginationRequest, totalItems int64) PaginationResponse {

	totalPages := int32(math.Ceil(float64(totalItems) / float64(req.PageSize)))

	return PaginationResponse{
		PageSize:    req.PageSize,
		CurrentPage: req.CurrentPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		Items:       items,
	}
}
