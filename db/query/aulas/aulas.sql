-- name: GetAulaByID :one
SELECT * FROM aulas WHERE id = $1;

-- name: ListAulas :many
SELECT * FROM aulas;

-- name: ListAulasByTimeRange :many
SELECT * FROM aulas WHERE
    meet_start >= $1 AND meet_end   <= $2 OR
    meet_end   >= $1 AND meet_start <= $2;

-- name: CreateAula :execresult
INSERT INTO aulas (tablet_id, professor_id, room_id, student_name, meet_start, meet_end, observation)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateAulaByID :execresult
UPDATE aulas SET room_id = $2, tablet_id = $3, professor_id = $4, student_name = $5 WHERE id = $1;

-- name: DeleteAulaByID :exec
DELETE FROM aulas WHERE id = 1;