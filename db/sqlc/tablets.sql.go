// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: tablets.sql

package sqlc

import (
	"context"
	"time"
)

const createTablet = `-- name: CreateTablet :one
INSERT INTO tablets (name, label_color) VALUES ($1, $2) RETURNING id, name, label_color, created_at, updated_at
`

type CreateTabletParams struct {
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

func (q *Queries) CreateTablet(ctx context.Context, arg CreateTabletParams) (Tablet, error) {
	row := q.db.QueryRowContext(ctx, createTablet, arg.Name, arg.LabelColor)
	var i Tablet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTabletByID = `-- name: DeleteTabletByID :exec
DELETE FROM tablets WHERE id = $1
`

func (q *Queries) DeleteTabletByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTabletByID, id)
	return err
}

const getTabletByID = `-- name: GetTabletByID :one
SELECT id, name, label_color, created_at, updated_at FROM tablets WHERE id = $1
`

func (q *Queries) GetTabletByID(ctx context.Context, id int32) (Tablet, error) {
	row := q.db.QueryRowContext(ctx, getTabletByID, id)
	var i Tablet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAvailableTabletsByTimeRange = `-- name: ListAvailableTabletsByTimeRange :many
SELECT id, name, label_color, created_at, updated_at FROM tablets WHERE id NOT IN (
     SELECT room_id FROM aulas a
     WHERE (meet_start >= $1 AND meet_end   <= $2 OR
            meet_end   >= $1 AND meet_start <= $2)
     AND room_id IS NOT NULL
)
`

type ListAvailableTabletsByTimeRangeParams struct {
	MeetStart time.Time `json:"meetStart"`
	MeetEnd   time.Time `json:"meetEnd"`
}

func (q *Queries) ListAvailableTabletsByTimeRange(ctx context.Context, arg ListAvailableTabletsByTimeRangeParams) ([]Tablet, error) {
	rows, err := q.db.QueryContext(ctx, listAvailableTabletsByTimeRange, arg.MeetStart, arg.MeetEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tablet{}
	for rows.Next() {
		var i Tablet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LabelColor,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTablets = `-- name: ListTablets :many
SELECT count(*) OVER () AS total_items, sub_query.id, sub_query.name, sub_query.label_color, sub_query.created_at, sub_query.updated_at FROM
    (SELECT id, name, label_color, created_at, updated_at FROM tablets ORDER BY
         CASE WHEN NOT  $3::bool AND $4::text = 'name' THEN name END ASC,
         CASE WHEN      $3::bool AND $4::text = 'name' THEN name END DESC,
         CASE WHEN NOT  $3::bool AND $4::text = 'id'   THEN id END ASC,
         CASE WHEN      $3::bool AND $4::text = 'id'   THEN id END DESC
    ) sub_query LIMIT $1 OFFSET $2
`

type ListTabletsParams struct {
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
	Reverse bool   `json:"reverse"`
	OrderBy string `json:"orderBy"`
}

type ListTabletsRow struct {
	TotalItems int64     `json:"totalItems"`
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	LabelColor string    `json:"labelColor"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (q *Queries) ListTablets(ctx context.Context, arg ListTabletsParams) ([]ListTabletsRow, error) {
	rows, err := q.db.QueryContext(ctx, listTablets,
		arg.Limit,
		arg.Offset,
		arg.Reverse,
		arg.OrderBy,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTabletsRow{}
	for rows.Next() {
		var i ListTabletsRow
		if err := rows.Scan(
			&i.TotalItems,
			&i.ID,
			&i.Name,
			&i.LabelColor,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTabletByID = `-- name: UpdateTabletByID :one
UPDATE tablets SET name = $2, label_color = $3 WHERE id = $1 RETURNING id, name, label_color, created_at, updated_at
`

type UpdateTabletByIDParams struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	LabelColor string `json:"labelColor"`
}

func (q *Queries) UpdateTabletByID(ctx context.Context, arg UpdateTabletByIDParams) (Tablet, error) {
	row := q.db.QueryRowContext(ctx, updateTabletByID, arg.ID, arg.Name, arg.LabelColor)
	var i Tablet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
