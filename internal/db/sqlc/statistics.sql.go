// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: statistics.sql

package db

import (
	"context"
)

const createStats = `-- name: CreateStats :one
INSERT INTO "statistics" (
  chat_id
) VALUES ($1) RETURNING id, chat_id, count, created_at
`

func (q *Queries) CreateStats(ctx context.Context, chatID string) (Statistic, error) {
	row := q.db.QueryRow(ctx, createStats, chatID)
	var i Statistic
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.Count,
		&i.CreatedAt,
	)
	return i, err
}

const getStats = `-- name: GetStats :one
SELECT id, chat_id, count, created_at FROM "statistics" WHERE chat_id = $1 LIMIT 1
`

func (q *Queries) GetStats(ctx context.Context, chatID string) (Statistic, error) {
	row := q.db.QueryRow(ctx, getStats, chatID)
	var i Statistic
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.Count,
		&i.CreatedAt,
	)
	return i, err
}

const updateCounter = `-- name: UpdateCounter :exec
UPDATE "statistics"
SET
  count = count + 1
WHERE
  chat_id = $1
`

func (q *Queries) UpdateCounter(ctx context.Context, chatID string) error {
	_, err := q.db.Exec(ctx, updateCounter, chatID)
	return err
}
