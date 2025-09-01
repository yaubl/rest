-- name: CreateBot :one
INSERT INTO bots (id, author, name, description) 
VALUES (?1, ?2, ?3, ?4)
RETURNING *;

-- name: GetBot :one
SELECT * FROM bots
WHERE id = ?1;

-- name: ListBots :many
SELECT * FROM bots
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: ListBotsByAuthor :many
SELECT * FROM bots
WHERE author = ?1
ORDER BY created_at DESC
LIMIT ?2 OFFSET ?3;

-- name: ListBotsByStatus :many
SELECT * FROM bots
WHERE status = ?1
ORDER BY created_at DESC
LIMIT ?2 OFFSET ?3;

-- name: UpdateBot :one
UPDATE bots
SET
    name        = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    status      = COALESCE(sqlc.narg('status'), status)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteBot :exec
DELETE FROM bots
WHERE id = ?1;

