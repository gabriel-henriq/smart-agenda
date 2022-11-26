-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = $1;

-- name: ListRooms :many
SELECT COUNT(*) OVER () AS total_items, sub_query.* FROM
    (SELECT * FROM  rooms ORDER BY name) sub_query LIMIT $1 OFFSET $2;

-- name: CreateRoom :one
INSERT INTO rooms (name) VALUES ($1) RETURNING *;

-- name: DeleteRoomByID :exec
DELETE FROM rooms WHERE id = $1;

-- name: UpdateRoomByID :execresult
UPDATE rooms SET name = $2 WHERE id = $1;

-- name: ListAvailableRoomsByTimeRange :many
SELECT *
FROM rooms
WHERE id NOT IN (SELECT room_id
                 FROM aulas a
                 WHERE (meet_start >= $1 AND meet_end   <= $2 OR
                        meet_end   >= $1 AND meet_start <= $2)
                   AND room_id IS NOT NULL);