-- name: CreateUser :one
INSERT INTO users (username)
VALUES (?)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

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

