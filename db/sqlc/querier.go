// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAula(ctx context.Context, arg CreateAulaParams) (sql.Result, error)
	CreateProfessor(ctx context.Context, arg CreateProfessorParams) (sql.Result, error)
	CreateRoom(ctx context.Context, name sql.NullString) (sql.Result, error)
	CreateTablet(ctx context.Context, name sql.NullString) (sql.Result, error)
	DeleteAulaByID(ctx context.Context) error
	DeleteProfessorByID(ctx context.Context, id int32) error
	DeleteRoomByID(ctx context.Context, id int32) error
	DeleteTabletByID(ctx context.Context, id int32) error
	GetAulaByID(ctx context.Context, id int32) (Aula, error)
	GetProfessorByID(ctx context.Context, id int32) (Professor, error)
	GetRoomByID(ctx context.Context, id int32) (Room, error)
	GetTabletByID(ctx context.Context, id int32) (Tablet, error)
	ListAulas(ctx context.Context) ([]Aula, error)
	ListAulasByTimeRange(ctx context.Context, arg ListAulasByTimeRangeParams) ([]Aula, error)
	ListAvailableProfessorsByTimeRange(ctx context.Context, arg ListAvailableProfessorsByTimeRangeParams) ([]Professor, error)
	ListAvailableRoomsByTimeRange(ctx context.Context, arg ListAvailableRoomsByTimeRangeParams) ([]Room, error)
	ListAvailableTabletsByTimeRange(ctx context.Context, arg ListAvailableTabletsByTimeRangeParams) ([]Tablet, error)
	ListProfessors(ctx context.Context) ([]Aula, error)
	ListRooms(ctx context.Context) ([]Room, error)
	ListTablets(ctx context.Context) ([]Tablet, error)
	UpdateAulaByID(ctx context.Context, arg UpdateAulaByIDParams) (sql.Result, error)
	UpdateProfessorByID(ctx context.Context, arg UpdateProfessorByIDParams) (sql.Result, error)
	UpdateRoomByID(ctx context.Context, arg UpdateRoomByIDParams) (sql.Result, error)
	UpdateTabletByID(ctx context.Context, arg UpdateTabletByIDParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
