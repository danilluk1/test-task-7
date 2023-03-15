package middlewares

import (
	"context"
	"time"

	"github.com/danilluk1/test-task-7/internal/types"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog/log"
)

type LoggMiddleware struct {
	Services *types.Services
}

func (c *LoggMiddleware) Wrap(next tgb.Handler) tgb.Handler {
	return tgb.HandlerFunc(func(ctx context.Context, update *tgb.Update) error {
		defer func(started time.Time) {
			log.Info().Dur("duration", time.Since(started)).Msg("update handled")
		}(time.Now())

		return next.Handle(ctx, update)
	})
}
