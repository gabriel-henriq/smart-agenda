-- name: GetTabletByID :one
SELECT * FROM tablets WHERE id = $1;

-- name: ListTablets :many
SELECT * FROM tablets;

-- name: UpdateTabletByID :execresult
UPDATE tablets SET name = $2 WHERE id = $1;

-- name: CreateTablet :execresult
INSERT INTO tablets (name) VALUES ($1);

-- name: DeleteTabletByID :exec
DELETE FROM tablets WHERE id = $1;

-- name: ListAvailableTabletsByTimeRange :many
SELECT *
FROM tablets
WHERE id NOT IN (SELECT tablet_id
                 FROM aulas a
                 WHERE (meet_start >= $1 AND meet_end   <= $2 OR
                        meet_end   >= $1 AND meet_start <= $2)
                   AND tablet_id IS NOT NULL);
