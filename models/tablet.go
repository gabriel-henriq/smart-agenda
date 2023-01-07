package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
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

type TabletResponse struct {
	ID        int32  `json:"ID"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type ListTabletsResponse struct {
	Tablets            []TabletResponse `json:"tablets"`
	PaginationResponse `json:"pagination"`
}

func TabletToJSON(sqlTablet sqlc.Tablet) TabletResponse {
	return TabletResponse{
		ID:        sqlTablet.ID,
		Name:      sqlTablet.Name,
		CreatedAt: sqlTablet.CreatedAt.Unix(),
		UpdatedAt: sqlTablet.UpdatedAt.Unix(),
	}
}

func TabletsToJSONList(SQLTablets []sqlc.ListTabletsRow) ListTabletsResponse {
	var tablets []TabletResponse

	for _, tablet := range SQLTablets {
		tablets = append(tablets, TabletResponse{
			ID:        tablet.ID,
			Name:      tablet.Name,
			CreatedAt: tablet.CreatedAt.Unix(),
			UpdatedAt: tablet.UpdatedAt.Unix(),
		})
	}

	return ListTabletsResponse{Tablets: tablets}
}
