-- name: CreateUser :one
INSERT INTO users (id, username)
VALUES (?, ?)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateBot :one
INSERT INTO bots (author, name, description)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetBotByID :one
SELECT * FROM bots
WHERE id = ?;

-- name: ListBotsByUserID :many
SELECT * FROM bots
WHERE author = ?
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET username = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: CountUsers :one
SELECT COUNT(*) AS count FROM users;

-- name: UpdateBot :one
UPDATE bots
SET name = ?, description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteBot :exec
DELETE FROM bots
WHERE id = ?;

-- name: CountBotsByUserID :one
SELECT COUNT(*) AS count
FROM bots
WHERE author = ?;

-- name: SearchBotsByName :many
SELECT * FROM bots
WHERE LOWER(name) LIKE LOWER('%' || ? || '%')
ORDER BY created_at DESC;
