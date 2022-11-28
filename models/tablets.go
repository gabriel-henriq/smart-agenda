package models

import (
	"github.com/gabriel-henriq/smart-agenda/db/sqlc"
)

type TabletList struct {
	Tablets    []sqlc.Tablet `json:"tablets"`
	Pagination `json:"pagination"`
}
