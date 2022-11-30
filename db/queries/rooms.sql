-- name: GetRoomByID :one
SELECT * FROM rooms WHERE id = $1;

-- name: ListRooms :many
SELECT count(*) OVER () AS total_items, sub_query.* FROM
    (SELECT * FROM rooms
     ORDER BY CASE
      WHEN NOT @reverse::bool AND @order_by::text = 'name' THEN name
      END ASC, CASE
                   WHEN @reverse::bool AND @order_by::text = 'name' THEN name
      END DESC, CASE
                    WHEN NOT @reverse::bool AND @order_by::text = 'id' THEN id
      END ASC, CASE
                   WHEN @reverse::bool AND @order_by::text = 'id' THEN id
      END DESC)
        sub_query LIMIT $1 OFFSET $2;

-- name: CreateRoom :one
INSERT INTO rooms (name, label_color) VALUES ($1, $2) RETURNING *;

-- name: DeleteRoomByID :exec
DELETE FROM rooms WHERE id = $1;

-- name: UpdateRoomByID :one
UPDATE rooms SET name = $2, label_color = $3 WHERE id = $1 RETURNING *;

-- name: ListAvailableRoomsByTimeRange :many
SELECT *
FROM rooms
WHERE id NOT IN (SELECT room_id
                 FROM aulas a
                 WHERE (meet_start >= $1 AND meet_end   <= $2 OR
                        meet_end   >= $1 AND meet_start <= $2)
                   AND room_id IS NOT NULL);
