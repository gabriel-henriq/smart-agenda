package tablets

import (
	"math"

	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
	"github.com/gabriel-henriq/smart-agenda/models"
)

func (r Tablet) toJSONTablet(sqlTablet sqlc.Tablet) models.Tablet {
	return models.Tablet{
		ID:        sqlTablet.ID,
		Name:      sqlTablet.Name.String,
		CreatedAt: sqlTablet.CreatedAt.Time,
		UpdatedAt: sqlTablet.UpdatedAt.Time,
	}
}

func (r Tablet) toJSONTabletList(SQLTablets []sqlc.ListTabletsRow, pageID, pageSize int32) models.TabletList {
	var tablets []models.Tablet

	for _, room := range SQLTablets {
		tablets = append(tablets, models.Tablet{
			Name:      room.Name.String,
			CreatedAt: room.CreatedAt.Time,
			UpdatedAt: room.UpdatedAt.Time,
		})
	}

	totalPages := int32(math.Ceil(float64(SQLTablets[0].TotalItems) / float64(pageSize)))

	return models.TabletList{
		Tablets: tablets,
		Pagination: models.Pagination{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLTablets[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
