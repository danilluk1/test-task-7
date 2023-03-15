package types

import db "github.com/danilluk1/test-task-7/internal/db/sqlc"

type Services struct {
	I18N      *i18n.I18N
	Statistic db.Store
}
