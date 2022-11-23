-- name: GetProfessorByID :one
SELECT * FROM professors WHERE id = $1;

-- name: ListProfessors :many
SELECT * FROM aulas;

-- name: UpdateProfessorByID :execresult
UPDATE professors SET name = $2, label_color = $3 WHERE id = $1;

-- name: CreateProfessor :execresult
INSERT INTO professors (name, label_color) VALUES ($1, $2);

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
