-- name: GetProfessorByID :one
SELECT * FROM professors WHERE id = $1;

-- name: ListProfessors :many
SELECT COUNT(*) OVER () AS total_items, sub_query.* FROM
    (SELECT * FROM  professors ORDER BY name) sub_query LIMIT $1 OFFSET $2;

-- name: UpdateProfessorByID :exec
UPDATE professors SET name = $2, label_color = $3 WHERE id = $1;

-- name: CreateProfessor :one
INSERT INTO professors (name, label_color) VALUES ($1, $2) RETURNING *;

-- name: DeleteProfessorByID :exec
DELETE FROM professors WHERE id = $1;

-- name: ListAvailableProfessorsByTimeRange :many
SELECT *
FROM professors
WHERE id NOT IN (SELECT professor_id
                 FROM aulas a
                 WHERE (meet_start >= $1 AND meet_end   <= $2 OR
                        meet_end   >= $1 AND meet_start <= $2)
                   AND professor_id IS NOT NULL);
