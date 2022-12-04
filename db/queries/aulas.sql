-- name: GetAulaByID :one
SELECT * FROM aulas WHERE id = $1;

-- name: ListAulas :many
SELECT * FROM aulas;

-- name: ListAulasByTimeRange :many
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
    meet_end   >= $1 AND meet_start <= $2;

-- name: CreateAula :one
INSERT INTO aulas (tablet_id, professor_id, room_id, student_name, meet_start, meet_end, observation)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdateAulaByID :one
UPDATE aulas
SET room_id = COALESCE(sqlc.narg('room_id'), room_id),
    tablet_id = COALESCE(sqlc.narg('tablet_id'), tablet_id),
    professor_id = COALESCE(sqlc.narg('professor_id'), professor_id),
    student_name = COALESCE(sqlc.narg('student_name'), student_name),
    meet_start = COALESCE(sqlc.narg('meet_start'), meet_start),
    meet_end = COALESCE(sqlc.narg('meet_end'), meet_end),
    observation = COALESCE(sqlc.narg('observation'), observation)
WHERE id = sqlc.arg('id') RETURNING *;

-- name: DeleteAulaByID :one
DELETE FROM aulas WHERE id = $1 RETURNING *;