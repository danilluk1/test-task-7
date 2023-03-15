-- GetStats :one
SELECT * FROM "statistics";

-- UpdateCounter :exec
UPDATE "statistics"
SET
  count = count + 1
WHERE
  user_id = $1;