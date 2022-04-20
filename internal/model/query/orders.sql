-- name: GetOrdersById :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: CreateOrders :one
INSERT INTO orders (
    customerid
) VALUES (
  $1
)
RETURNING *;

-- name: ListOrders :many
SELECT * FROM orders
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteOrderById :exec
DELETE FROM orders WHERE id = $1;