-- name: CreateStats :one
INSERT INTO "statistics" (
  chat_id
) VALUES ($1) RETURNING *;

-- name: GetStats :one
SELECT * FROM "statistics" WHERE chat_id = $1 LIMIT 1;

-- name: UpdateCounter :exec
UPDATE "statistics"
SET
  count = count + 1
WHERE
  chat_id = $1;