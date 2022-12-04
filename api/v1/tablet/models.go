package tablet

import (
	"github.com/gabriel-henriq/smart-agenda/api/v1"
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"math"
)

type createRequest struct {
	Name       string `json:"name" binding:"required"`
	LabelColor string `json:"labelColor" binding:"required"`
}

type deleteRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type updateRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

type response struct {
	ID        int32  `json:"ID"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type listResponse struct {
	Tablets               []response `json:"tablets"`
	v1.PaginationResponse `json:"pagination"`
}

func toJSON(sqlTablet sqlc.Tablet) response {
	return response{
		ID:        sqlTablet.ID,
		Name:      sqlTablet.Name,
		CreatedAt: sqlTablet.CreatedAt.Unix(),
		UpdatedAt: sqlTablet.UpdatedAt.Unix(),
	}
}

func toJSONList(SQLTablets []sqlc.ListTabletsRow, pageID, pageSize int32) listResponse {
	var tablets []response

	for _, tablet := range SQLTablets {
		tablets = append(tablets, response{
			ID:        tablet.ID,
			Name:      tablet.Name,
			CreatedAt: tablet.CreatedAt.Unix(),
			UpdatedAt: tablet.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLTablets[0].TotalItems) / float64(pageSize)))

	return listResponse{
		Tablets: tablets,
		PaginationResponse: v1.PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLTablets[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
