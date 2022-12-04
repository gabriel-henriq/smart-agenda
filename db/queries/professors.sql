-- name: GetProfessorByID :one
SELECT * FROM professors WHERE id = $1;

-- name: ListProfessors :many
SELECT count(*) OVER () AS total_items, sub_query.* FROM
    (SELECT * FROM professors ORDER BY CASE
        WHEN NOT @reverse::bool AND @order_by::text = 'name' THEN name END ASC, CASE
        WHEN @reverse::bool     AND @order_by::text = 'name' THEN name END DESC, CASE
        WHEN NOT @reverse::bool AND @order_by::text = 'id'   THEN id   END ASC, CASE
        WHEN @reverse::bool     AND @order_by::text = 'id'   THEN id   END DESC)
        sub_query LIMIT $1 OFFSET $2;

-- name: CreateProfessor :one
INSERT INTO professors (name, label_color) VALUES ($1, $2) RETURNING *;

-- name: UpdateProfessorByID :one
UPDATE professors
SET
    name = COALESCE(sqlc.narg('name'), name),
    label_color = COALESCE(sqlc.narg('label_color'), label_color)
WHERE id = sqlc.arg('id') RETURNING *;

-- name: DeleteProfessorByID :one
DELETE FROM professors WHERE id = $1 RETURNING *;

-- name: ListAvailableProfessorsByTimeRange :many
SELECT *
FROM professors
WHERE id NOT IN (SELECT professor_id
     FROM aulas a
     WHERE (meet_start >= $1 AND meet_end   <= $2 OR
            meet_end   >= $1 AND meet_start <= $2)
     AND professor_id IS NOT NULL);
