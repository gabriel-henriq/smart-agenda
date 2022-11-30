// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: aulas.sql

package sqlc

import (
	"context"
	"time"
)

const createAula = `-- name: CreateAula :one
INSERT INTO aulas (tablet_id, professor_id, room_id, student_name, meet_start, meet_end, observation)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, tablet_id, professor_id, room_id, student_name, observation, meet_start, meet_end, created_at, updated_at
`

type CreateAulaParams struct {
	TabletID    int32     `json:"tabletID"`
	ProfessorID int32     `json:"professorID"`
	RoomID      int32     `json:"roomID"`
	StudentName string    `json:"studentName"`
	MeetStart   time.Time `json:"meetStart"`
	MeetEnd     time.Time `json:"meetEnd"`
	Observation string    `json:"observation"`
}

func (q *Queries) CreateAula(ctx context.Context, arg CreateAulaParams) (Aula, error) {
	row := q.db.QueryRowContext(ctx, createAula,
		arg.TabletID,
		arg.ProfessorID,
		arg.RoomID,
		arg.StudentName,
		arg.MeetStart,
		arg.MeetEnd,
		arg.Observation,
	)
	var i Aula
	err := row.Scan(
		&i.ID,
		&i.TabletID,
		&i.ProfessorID,
		&i.RoomID,
		&i.StudentName,
		&i.Observation,
		&i.MeetStart,
		&i.MeetEnd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAulaByID = `-- name: DeleteAulaByID :exec
DELETE FROM aulas WHERE id = $1
`

func (q *Queries) DeleteAulaByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAulaByID, id)
	return err
}

const getAulaByID = `-- name: GetAulaByID :one
SELECT id, tablet_id, professor_id, room_id, student_name, observation, meet_start, meet_end, created_at, updated_at FROM aulas WHERE id = $1
`

func (q *Queries) GetAulaByID(ctx context.Context, id int32) (Aula, error) {
	row := q.db.QueryRowContext(ctx, getAulaByID, id)
	var i Aula
	err := row.Scan(
		&i.ID,
		&i.TabletID,
		&i.ProfessorID,
		&i.RoomID,
		&i.StudentName,
		&i.Observation,
		&i.MeetStart,
		&i.MeetEnd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAulas = `-- name: ListAulas :many
SELECT id, tablet_id, professor_id, room_id, student_name, observation, meet_start, meet_end, created_at, updated_at FROM aulas
`

func (q *Queries) ListAulas(ctx context.Context) ([]Aula, error) {
	rows, err := q.db.QueryContext(ctx, listAulas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Aula{}
	for rows.Next() {
		var i Aula
		if err := rows.Scan(
			&i.ID,
			&i.TabletID,
			&i.ProfessorID,
			&i.RoomID,
			&i.StudentName,
			&i.Observation,
			&i.MeetStart,
			&i.MeetEnd,
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

const listAulasByTimeRange = `-- name: ListAulasByTimeRange :many
SELECT a.id,
       a.student_name,
       a.meet_start,
       a.meet_end,
       p.name AS professor_name,
       t.name AS tablet_name,
       r.name AS room_name,
       a.created_at,
       a.updated_at FROM aulas a
    JOIN professors p on p.id = a.professor_id
    JOIN tablets    t on t.id = a.tablet_id
    JOIN rooms      r on r.id = a.room_id
WHERE
    meet_start >= $1 AND meet_end   <= $2 OR
    meet_end   >= $1 AND meet_start <= $2
`

type ListAulasByTimeRangeParams struct {
	MeetStart time.Time `json:"meetStart"`
	MeetEnd   time.Time `json:"meetEnd"`
}

type ListAulasByTimeRangeRow struct {
	ID            int32     `json:"id"`
	StudentName   string    `json:"studentName"`
	MeetStart     time.Time `json:"meetStart"`
	MeetEnd       time.Time `json:"meetEnd"`
	ProfessorName string    `json:"professorName"`
	TabletName    string    `json:"tabletName"`
	RoomName      string    `json:"roomName"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (q *Queries) ListAulasByTimeRange(ctx context.Context, arg ListAulasByTimeRangeParams) ([]ListAulasByTimeRangeRow, error) {
	rows, err := q.db.QueryContext(ctx, listAulasByTimeRange, arg.MeetStart, arg.MeetEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAulasByTimeRangeRow{}
	for rows.Next() {
		var i ListAulasByTimeRangeRow
		if err := rows.Scan(
			&i.ID,
			&i.StudentName,
			&i.MeetStart,
			&i.MeetEnd,
			&i.ProfessorName,
			&i.TabletName,
			&i.RoomName,
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

const updateAulaByID = `-- name: UpdateAulaByID :one
UPDATE aulas SET room_id = $2, tablet_id = $3, professor_id = $4, student_name = $5, meet_start = $6, meet_end = $7, observation = $8 WHERE id = $1 RETURNING id, tablet_id, professor_id, room_id, student_name, observation, meet_start, meet_end, created_at, updated_at
`

type UpdateAulaByIDParams struct {
	ID          int32     `json:"id"`
	RoomID      int32     `json:"roomID"`
	TabletID    int32     `json:"tabletID"`
	ProfessorID int32     `json:"professorID"`
	StudentName string    `json:"studentName"`
	MeetStart   time.Time `json:"meetStart"`
	MeetEnd     time.Time `json:"meetEnd"`
	Observation string    `json:"observation"`
}

func (q *Queries) UpdateAulaByID(ctx context.Context, arg UpdateAulaByIDParams) (Aula, error) {
	row := q.db.QueryRowContext(ctx, updateAulaByID,
		arg.ID,
		arg.RoomID,
		arg.TabletID,
		arg.ProfessorID,
		arg.StudentName,
		arg.MeetStart,
		arg.MeetEnd,
		arg.Observation,
	)
	var i Aula
	err := row.Scan(
		&i.ID,
		&i.TabletID,
		&i.ProfessorID,
		&i.RoomID,
		&i.StudentName,
		&i.Observation,
		&i.MeetStart,
		&i.MeetEnd,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
