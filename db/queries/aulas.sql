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
UPDATE aulas SET room_id = $2, tablet_id = $3, professor_id = $4, student_name = $5, meet_start = $6, meet_end = $7, observation = $8 WHERE id = $1 RETURNING *;

-- name: DeleteAulaByID :exec
DELETE FROM aulas WHERE id = $1;