-- name: CreateOrderDetail :one
INSERT INTO order_detail (
    order_id,product_id,active,total
) VALUES (
  $1,$2,$3,$4
)
RETURNING *;
-- name: GetOrderDetail :one
SELECT * FROM order_detail
WHERE id = $1 LIMIT 1;
-- name: DeleteOrderDetailById :exec
DELETE FROM order_detail WHERE id = $1;
-- name: ListOrderDetail :many
SELECT * FROM order_detail
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;