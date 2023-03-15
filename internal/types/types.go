package types

import db "github.com/danilluk1/test-task-7/internal/db/sqlc"

type Services struct {
	Statistic db.Store
}
