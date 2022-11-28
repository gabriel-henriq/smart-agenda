package models

import "github.com/gabriel-henriq/smart-agenda/db/sqlc"

type ProfessorList struct {
	Professors []sqlc.Professor `json:"professors"`
	Pagination `json:"pagination"`
}
