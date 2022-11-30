-- name: GetTabletByID :one
SELECT * FROM tablets WHERE id = $1;

-- name: ListTablets :many
SELECT count(*) OVER () AS total_items, sub_query.* FROM
    (SELECT * FROM tablets ORDER BY
         CASE WHEN NOT  @reverse::bool AND @order_by::text = 'name' THEN name END ASC,
         CASE WHEN      @reverse::bool AND @order_by::text = 'name' THEN name END DESC,
         CASE WHEN NOT  @reverse::bool AND @order_by::text = 'id'   THEN id END ASC,
         CASE WHEN      @reverse::bool AND @order_by::text = 'id'   THEN id END DESC
    ) sub_query LIMIT $1 OFFSET $2;

-- name: CreateTablet :one
INSERT INTO tablets (name, label_color) VALUES ($1, $2) RETURNING *;

-- name: DeleteTabletByID :exec
DELETE FROM tablets WHERE id = $1;

-- name: UpdateTabletByID :one
UPDATE tablets SET name = $2, label_color = $3 WHERE id = $1 RETURNING *;

-- name: ListAvailableTabletsByTimeRange :many
SELECT * FROM tablets WHERE id NOT IN (
     SELECT room_id FROM aulas a
     WHERE (meet_start >= $1 AND meet_end   <= $2 OR
            meet_end   >= $1 AND meet_start <= $2)
     AND room_id IS NOT NULL
);
