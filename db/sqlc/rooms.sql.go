// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: rooms.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (name, label_color) VALUES ($1, $2) RETURNING id, name, label_color, created_at, updated_at
`

type CreateRoomParams struct {
	Name       sql.NullString `json:"name"`
	LabelColor sql.NullString `json:"labelColor"`
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

const deleteRoomByID = `-- name: DeleteRoomByID :exec
DELETE FROM rooms WHERE id = $1
`

func (q *Queries) DeleteRoomByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRoomByID, id)
	return err
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
	MeetStart sql.NullTime `json:"meetStart"`
	MeetEnd   sql.NullTime `json:"meetEnd"`
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
    (SELECT id, name, label_color, created_at, updated_at FROM rooms) sub_query LIMIT $1 OFFSET $2
`

type ListRoomsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListRoomsRow struct {
	TotalItems int64          `json:"totalItems"`
	ID         int32          `json:"id"`
	Name       sql.NullString `json:"name"`
	LabelColor sql.NullString `json:"labelColor"`
	CreatedAt  sql.NullTime   `json:"createdAt"`
	UpdatedAt  sql.NullTime   `json:"updatedAt"`
}

func (q *Queries) ListRooms(ctx context.Context, arg ListRoomsParams) ([]ListRoomsRow, error) {
	rows, err := q.db.QueryContext(ctx, listRooms, arg.Limit, arg.Offset)
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

const updateRoomByID = `-- name: UpdateRoomByID :execresult
UPDATE rooms SET name = $2 WHERE id = $1
`

type UpdateRoomByIDParams struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
}

func (q *Queries) UpdateRoomByID(ctx context.Context, arg UpdateRoomByIDParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateRoomByID, arg.ID, arg.Name)
}
