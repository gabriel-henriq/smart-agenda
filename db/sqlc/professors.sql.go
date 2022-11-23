// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: professors.sql

package db

import (
	"context"
	"database/sql"
)

const createProfessor = `-- name: CreateProfessor :execresult
INSERT INTO professors (name, label_color) VALUES ($1, $2)
`

type CreateProfessorParams struct {
	Name       sql.NullString `json:"name"`
	LabelColor sql.NullString `json:"labelColor"`
}

func (q *Queries) CreateProfessor(ctx context.Context, arg CreateProfessorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProfessor, arg.Name, arg.LabelColor)
}

const deleteProfessorByID = `-- name: DeleteProfessorByID :exec
DELETE FROM professors WHERE id = $1
`

func (q *Queries) DeleteProfessorByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProfessorByID, id)
	return err
}

const getProfessorByID = `-- name: GetProfessorByID :one
SELECT id, name, label_color, created_at, updated_at FROM professors WHERE id = $1
`

func (q *Queries) GetProfessorByID(ctx context.Context, id int32) (Professor, error) {
	row := q.db.QueryRowContext(ctx, getProfessorByID, id)
	var i Professor
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LabelColor,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAvailableProfessorsByTimeRange = `-- name: ListAvailableProfessorsByTimeRange :many
SELECT id, name, label_color, created_at, updated_at
FROM professors
WHERE id NOT IN (SELECT professor_id
                 FROM aulas a
                 WHERE (meet_start >= $1 AND meet_end   <= $2 OR
                        meet_end   >= $1 AND meet_start <= $2)
                   AND professor_id IS NOT NULL)
`

type ListAvailableProfessorsByTimeRangeParams struct {
	MeetStart sql.NullTime `json:"meetStart"`
	MeetEnd   sql.NullTime `json:"meetEnd"`
}

func (q *Queries) ListAvailableProfessorsByTimeRange(ctx context.Context, arg ListAvailableProfessorsByTimeRangeParams) ([]Professor, error) {
	rows, err := q.db.QueryContext(ctx, listAvailableProfessorsByTimeRange, arg.MeetStart, arg.MeetEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Professor{}
	for rows.Next() {
		var i Professor
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

const listProfessors = `-- name: ListProfessors :many
SELECT id, tablet_id, professor_id, room_id, student_name, meet_start, meet_end, observation, created_at, updated_at FROM aulas
`

func (q *Queries) ListProfessors(ctx context.Context) ([]Aula, error) {
	rows, err := q.db.QueryContext(ctx, listProfessors)
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
			&i.MeetStart,
			&i.MeetEnd,
			&i.Observation,
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

const updateProfessorByID = `-- name: UpdateProfessorByID :execresult
UPDATE professors SET name = $2, label_color = $3 WHERE id = $1
`

type UpdateProfessorByIDParams struct {
	ID         int32          `json:"id"`
	Name       sql.NullString `json:"name"`
	LabelColor sql.NullString `json:"labelColor"`
}

func (q *Queries) UpdateProfessorByID(ctx context.Context, arg UpdateProfessorByIDParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProfessorByID, arg.ID, arg.Name, arg.LabelColor)
}
