-- name: GetProductById :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;


