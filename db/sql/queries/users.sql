-- name: CreateUser :one
INSERT INTO users (id, username) 
VALUES (?1, ?2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: UpdateUsername :one
UPDATE users
SET username = ?2
WHERE id = ?1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?1;

