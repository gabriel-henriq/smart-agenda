// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: rooms.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (name, label_color) VALUES ($1, $2) RETURNING id, name, label_color, created_at, updated_at
`

type CreateRoomParams struct {
	Name       string
	LabelColor string
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error) {
	row := q.db.QueryRowContext(ctx, createRoom, arg.Name, arg.LabelColor)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRoomByID = `-- name: DeleteRoomByID :one
DELETE FROM rooms WHERE id = $1 RETURNING id, name, label_color, created_at, updated_at
`

func (q *Queries) DeleteRoomByID(ctx context.Context, id int32) (Room, error) {
	row := q.db.QueryRowContext(ctx, deleteRoomByID, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoomByID = `-- name: GetRoomByID :one
SELECT id, name, label_color, created_at, updated_at FROM rooms WHERE id = $1
`

func (q *Queries) GetRoomByID(ctx context.Context, id int32) (Room, error) {
	row := q.db.QueryRowContext(ctx, getRoomByID, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAvailableRoomsByTimeRange = `-- name: ListAvailableRoomsByTimeRange :many
SELECT id, name, label_color, created_at, updated_at
FROM rooms
WHERE id NOT IN (SELECT room_id
     FROM aulas a
     WHERE (meet_start >= $1 AND meet_end   <= $2 OR
            meet_end   >= $1 AND meet_start <= $2)
       AND room_id IS NOT NULL)
`

type ListAvailableRoomsByTimeRangeParams struct {
	MeetStart time.Time
	MeetEnd   time.Time
}

func (q *Queries) ListAvailableRoomsByTimeRange(ctx context.Context, arg ListAvailableRoomsByTimeRangeParams) ([]Room, error) {
	rows, err := q.db.QueryContext(ctx, listAvailableRoomsByTimeRange, arg.MeetStart, arg.MeetEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
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

const listRooms = `-- name: ListRooms :many
SELECT count(*) OVER () AS total_items, sub_query.id, sub_query.name, sub_query.label_color, sub_query.created_at, sub_query.updated_at FROM
    (SELECT id, name, label_color, created_at, updated_at FROM rooms
     ORDER BY CASE
      WHEN NOT $3::bool AND $4::text = 'name' THEN name END ASC, CASE
      WHEN $3::bool     AND $4::text = 'name' THEN name END DESC, CASE
      WHEN NOT $3::bool AND $4::text = 'id'   THEN id   END ASC, CASE
      WHEN $3::bool     AND $4::text = 'id'   THEN id   END DESC)
        sub_query LIMIT $1 OFFSET $2
`

type ListRoomsParams struct {
	Limit   int32
	Offset  int32
	Reverse bool
	OrderBy string
}

type ListRoomsRow struct {
	TotalItems int64
	ID         int32
	Name       string
	LabelColor string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (q *Queries) ListRooms(ctx context.Context, arg ListRoomsParams) ([]ListRoomsRow, error) {
	rows, err := q.db.QueryContext(ctx, listRooms,
		arg.Limit,
		arg.Offset,
		arg.Reverse,
		arg.OrderBy,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRoomsRow{}
	for rows.Next() {
		var i ListRoomsRow
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

const updateRoomByID = `-- name: UpdateRoomByID :one
UPDATE rooms
SET
    name = COALESCE($1, name),
    label_color = coalesce($2, label_color)
WHERE id = $3 RETURNING id, name, label_color, created_at, updated_at
`

type UpdateRoomByIDParams struct {
	Name       sql.NullString
	LabelColor sql.NullString
	ID         int32
}

func (q *Queries) UpdateRoomByID(ctx context.Context, arg UpdateRoomByIDParams) (Room, error) {
	row := q.db.QueryRowContext(ctx, updateRoomByID, arg.Name, arg.LabelColor, arg.ID)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
