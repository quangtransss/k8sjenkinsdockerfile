-- name: CreateRolePermisstion :one
INSERT INTO role_permisstion (
    roleid
) VALUES (
  $1
)
RETURNING *;