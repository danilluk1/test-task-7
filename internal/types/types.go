package types

import (
	db "github.com/danilluk1/test-task-7/internal/db/sqlc"
	"github.com/danilluk1/test-task-7/internal/weather"
)

type Services struct {
	Store          db.Store
	WeatherService *weather.Service
}
