-- name: CreateRole :one
INSERT INTO role (
    title,
    slug,
    active,
    description
) VALUES (
  $1, $2, $3 ,$4
)
RETURNING *;
-- name: ListRole :many
SELECT * FROM role
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;
-- name: GetRole :one
SELECT * FROM role
WHERE id = $1 LIMIT 1;