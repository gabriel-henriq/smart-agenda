package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type CreateTabletRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type DeleteTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetTabletRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type UpdateTabletRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type ResponseTablet struct {
	ID        int32  `json:"ID"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type ListTabletsResponse struct {
	Tablets            []ResponseTablet `json:"tablets"`
	PaginationResponse `json:"pagination"`
}

func TabletToJSON(sqlTablet sqlc.Tablet) ResponseTablet {
	return ResponseTablet{
		ID:        sqlTablet.ID,
		Name:      sqlTablet.Name,
		CreatedAt: sqlTablet.CreatedAt.Unix(),
		UpdatedAt: sqlTablet.UpdatedAt.Unix(),
	}
}

func TabletsToJSONList(SQLTablets []sqlc.ListTabletsRow, pageID, pageSize int32) ListTabletsResponse {
	var tablets []ResponseTablet

	for _, tablet := range SQLTablets {
		tablets = append(tablets, ResponseTablet{
			ID:        tablet.ID,
			Name:      tablet.Name,
			CreatedAt: tablet.CreatedAt.Unix(),
			UpdatedAt: tablet.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLTablets[0].TotalItems) / float64(pageSize)))

	return ListTabletsResponse{
		Tablets: tablets,
		PaginationResponse: PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLTablets[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
