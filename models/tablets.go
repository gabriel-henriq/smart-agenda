package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type CreateTabletRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type UpdateTabletRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type TabletResponse struct {
	ID        int32  `json:"ID"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type DeleteTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type TabletList struct {
	Tablets    []TabletResponse `json:"tablets"`
	Pagination `json:"pagination"`
}

type ListTabletRequest struct {
	PageSize int32  `form:"page_size"`
	PageID   int32  `form:"page_id"`
	OrderBy  string `form:"order_by"`
	Reverse  bool   `form:"reverse"`
}

func ToJSONTablet(sqlTablet sqlc.Tablet) TabletResponse {
	return TabletResponse{
		ID:        sqlTablet.ID,
		Name:      sqlTablet.Name,
		CreatedAt: sqlTablet.CreatedAt.String(),
		UpdatedAt: sqlTablet.UpdatedAt.String(),
	}
}

func ToJSONTabletList(SQLTablets []sqlc.ListTabletsRow, pageID, pageSize int32) TabletList {
	var tablets []TabletResponse

	for _, room := range SQLTablets {
		tablets = append(tablets, TabletResponse{
			ID:        room.ID,
			Name:      room.Name,
			CreatedAt: room.CreatedAt.String(),
			UpdatedAt: room.UpdatedAt.String(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLTablets[0].TotalItems) / float64(pageSize)))

	return TabletList{
		Tablets: tablets,
		Pagination: Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLTablets[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
