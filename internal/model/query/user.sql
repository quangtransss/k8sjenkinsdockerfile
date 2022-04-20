-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
-- name: CreateUser :one
INSERT INTO users (
  full_name,
  username,
  hashed_password,
  email,
  mobile,
  roleid
) VALUES (
  $1, $2, $3, $4 ,$5 ,$6
)
RETURNING *;
-- name: ListUsers :many
SELECT * FROM users
WHERE username = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;