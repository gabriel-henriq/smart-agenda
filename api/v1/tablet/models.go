package tablet

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
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
	Tablets                   []TabletResponse `json:"tablets"`
	models.PaginationResponse `json:"pagination"`
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

	for _, tablet := range SQLTablets {
		tablets = append(tablets, TabletResponse{
			ID:        tablet.ID,
			Name:      tablet.Name,
			CreatedAt: tablet.CreatedAt.String(),
			UpdatedAt: tablet.UpdatedAt.String(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLTablets[0].TotalItems) / float64(pageSize)))

	return TabletList{
		Tablets: tablets,
		PaginationResponse: models.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLTablets[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
