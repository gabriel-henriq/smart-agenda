-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = $1;

-- name: ListRooms :many
SELECT * FROM rooms;

-- name: CreateRoom :execresult
INSERT INTO rooms (name) VALUES ($1);

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